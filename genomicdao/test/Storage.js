const {
  time,
  loadFixture
} = require("@nomicfoundation/hardhat-toolbox/network-helpers");

const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("Controller", function () {
  async function deployControllerFixture() {
    const [owner, addr1, addr2] = await ethers.getSigners();

    const nft = await ethers.deployContract("GeneNFT");
    const pcspToken = await ethers.deployContract("PostCovidStrokePrevention");

    const controller = await ethers.deployContract("Controller", [nft.target, pcspToken.target]);

    await nft.transferOwnership(controller.target)
    await pcspToken.transferOwnership(controller.target)

    return { controller, nft, pcspToken, owner, addr1, addr2 }
  }

  describe("Upload Data", function () {
    it("Should receive session id", async function () {
      const { controller } = await loadFixture(deployControllerFixture);

      await expect(
        controller.uploadData("doc1")
      )
        .to.emit(controller, "UploadData")
        .withArgs("doc1", 0)
    })

    it("Should fail if the doc is submited", async function () {
      const { controller, addr1, addr2 } = await loadFixture(deployControllerFixture);

      const docId = "doc1"
      const contentHash = "dochash"
      const proof = "success"
      const riskScore = 1
      const sessionId = 0

      await controller.connect(addr1).uploadData(docId)
      await controller.connect(addr1).confirm(docId, contentHash, proof, sessionId, riskScore)

      await expect(
        controller.connect(addr2).uploadData(docId)
      ).to.be.revertedWith("Doc already been submitted")
    })
  })

  describe("Confirm data", function () {
    it("Should receive correct nft", async function () {
      const { controller, nft, owner } = await loadFixture(deployControllerFixture);

      const docId = "doc1"
      const contentHash = "dochash"
      const proof = "success"
      const riskScore = 1
      const sessionId = 0

      await controller.uploadData(docId)
      await controller.confirm(docId, contentHash, proof, sessionId, riskScore)

      expect(await nft.ownerOf(0)).to.equal(owner.address);
    })

    it("Should receive correct pcsp reward", async function () {
      const { controller, pcspToken, addr1 } = await loadFixture(deployControllerFixture);

      const docId = "doc1"
      const contentHash = "dochash"
      const proof = "success"
      const riskScore = 1
      const sessionId = 0

      const awardAmount = BigInt("15000") * BigInt("10") ** BigInt("18")

      await controller.connect(addr1).uploadData(docId)
      await controller.connect(addr1).confirm(docId, contentHash, proof, sessionId, riskScore)

      const ownerBalance = await pcspToken.balanceOf(addr1.address)

      expect(ownerBalance).to.equal(awardAmount)
    })

    it("Should close session", async function () {
      const { controller, addr1 } = await loadFixture(deployControllerFixture);

      const docId = "doc1"
      const contentHash = "dochash"
      const proof = "success"
      const riskScore = 1
      const sessionId = 0

      await controller.connect(addr1).uploadData(docId)
      await controller.connect(addr1).confirm(docId, contentHash, proof, sessionId, riskScore)

      const session = await controller.getSession(sessionId)

      expect(session.proof).to.equal(proof)
      expect(session.confirmed).to.equal(true)
    })

    it("Should content hash uploaded", async function () {
      const { controller, addr1 } = await loadFixture(deployControllerFixture);

      const docId = "doc1"
      const contentHash = "dochash"
      const proof = "success"
      const riskScore = 1
      const sessionId = 0

      await controller.connect(addr1).uploadData(docId)
      await controller.connect(addr1).confirm(docId, contentHash, proof, sessionId, riskScore)

      const doc = await controller.getDoc(docId)

      expect(doc.hashContent).to.equal(contentHash)
    })

    it("Should fail if the doc is submitted", async function () {
      const { controller, addr1 } = await loadFixture(deployControllerFixture);

      const docId = "doc1"
      const contentHash = "dochash"
      const proof = "success"
      const riskScore = 1
      const sessionId = 0

      await controller.connect(addr1).uploadData(docId)
      await controller.connect(addr1).confirm(docId, contentHash, proof, sessionId, riskScore)

      await expect(
        controller.connect(addr1).confirm(docId, contentHash, proof, sessionId, riskScore)
      ).to.be.revertedWith("Doc already been submitted")
    })

    it("Should fail if the session owner is invalid", async function () {
      const { controller, addr1, addr2 } = await loadFixture(deployControllerFixture);

      const docId = "doc1"
      const contentHash = "dochash"
      const proof = "success"
      const riskScore = 1
      const sessionId = 0

      await controller.connect(addr1).uploadData(docId)

      await expect(
        controller.connect(addr2).confirm(docId, contentHash, proof, sessionId, riskScore)
      ).to.be.revertedWith("Invalid session owner")
    })

    it("Should fail if the session is end", async function () {
      const { controller, addr1 } = await loadFixture(deployControllerFixture);

      const docId = "doc1"
      const docId2 = "doc2"
      const contentHash = "dochash"
      const proof = "success"
      const riskScore = 1
      const sessionId = 0

      await controller.connect(addr1).uploadData(docId)
      await controller.connect(addr1).confirm(docId, contentHash, proof, sessionId, riskScore)

      await expect(
        controller.connect(addr1).confirm(docId2, contentHash, proof, sessionId, riskScore)
      ).to.be.revertedWith("Session is ended")
    })
  })

  async function deployControllerThenConfirm() {
    const {controller, addr1, addr2} = await loadFixture(deployControllerFixture);

    const docId = "doc1"
    const contentHash = "dochash"
    const proof = "success"
    const riskScore = 1
    const sessionId = 0

    await controller.uploadData(docId)
    const nft = await controller.confirm(docId, contentHash, proof, sessionId, riskScore)
    const nftId = nft.value
    return {controller, addr1, addr2, nftId, docId, contentHash}
  }

  describe("Grant Access", function () {
    it("should allow NFT owner to grant access to another address and emit GrantedAccess event", async function () {
      const { controller, addr1, nftId, docId, contentHash } = await loadFixture(deployControllerThenConfirm);


      await controller.grantAccess(nftId, addr1.address);

      const hasAccess = await controller.connect(addr1).requestAccess(nftId);
      expect(hasAccess.length).to.be.equal(2);
      expect(hasAccess[0]).to.be.equal(docId);
      expect(hasAccess[1]).to.be.equal(contentHash);

      await expect(controller.grantAccess(nftId, addr1.address))
        .to.emit(controller, "GrantedAccess")
        .withArgs(nftId, addr1.address);
    });

    it("should not allow non-owner to grant access", async function () {
      const {controller, addr1, nftId, addr2} = await loadFixture(deployControllerThenConfirm);

      await expect(
        controller.connect(addr1).grantAccess(nftId, addr2.address)
      ).to.be.revertedWith("You are not the owner of this NFT.");
    });
  });

  describe("Revoke Access", function () {
    it("should allow NFT owner to revoke access from another address and emit RevokedAccess event", async function () {
      const { controller, addr1, nftId, docId, contentHash } = await loadFixture(deployControllerThenConfirm);

      // step 1: Grant Access -> check requestAccess and event
      await controller.grantAccess(nftId, addr1.address);

      const dataDoc = await controller.connect(addr1).requestAccess(nftId);
      expect(dataDoc.length).to.be.equal(2);
      expect(dataDoc[0]).to.be.equal(docId);
      expect(dataDoc[1]).to.be.equal(contentHash);

      await expect(controller.grantAccess(nftId, addr1.address))
        .to.emit(controller, "GrantedAccess")
        .withArgs(nftId, addr1.address);

      // step 2: Revoke Access -> check event and requestAccess
      await controller.revokeAccess(nftId, addr1.address);

      await expect(controller.revokeAccess(nftId, addr1.address))
        .to.emit(controller, "RevokedAccess")
        .withArgs(nftId, addr1.address);

      await expect(
        controller.connect(addr1).requestAccess(nftId)
      ).to.be.revertedWith("Access denied.");
    });

    it("should not allow non-owner to revoke access", async function () {
      const {controller, addr1, nftId, addr2} = await loadFixture(deployControllerThenConfirm);

      await expect(
        controller.connect(addr1).revokeAccess(nftId, addr2.address)
      ).to.be.revertedWith("You are not the owner of this NFT.");
    });
  });

})