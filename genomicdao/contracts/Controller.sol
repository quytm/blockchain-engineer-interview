// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/utils/Counters.sol";
import "./NFT.sol";
import "./Token.sol";

contract Controller {
    using Counters for Counters.Counter;

    //
    // STATE VARIABLES
    //
    Counters.Counter private _sessionIdCounter;
    GeneNFT public geneNFT;
    PostCovidStrokePrevention public pcspToken;

    struct UploadSession {
        uint256 id;
        address user;
        string proof;
        bool confirmed;
    }

    struct DataDoc {
        string id;
        string hashContent;
    }

    mapping(uint256 => UploadSession) sessions;
    mapping(string => DataDoc) docs;
    mapping(string => bool) docSubmits;
    mapping(uint256 => string) nftDocs;

    //
    // EVENTS
    //
    // UploadData: Gene data is submitted
    event UploadData(string docId, uint256 sessionId);
    // MintedGNFT: GNFT is minted
    event MintedGNFT(string docId, uint256 sessionId, uint256 tokenId);
    // RewardedPCSP: PCSP is rewarded
    event RewardedPCSP(string docId, uint256 sessionId, uint256 riskScore);

    constructor(address nftAddress, address pcspAddress) {
        geneNFT = GeneNFT(nftAddress);
        pcspToken = PostCovidStrokePrevention(pcspAddress);
    }

    function uploadData(string memory docId) public returns (uint256) {
        // 1. Check if the doc id has been submitted to the system before.
        require(!docSubmits[docId], "Doc already been submitted");

        // 2. Increase sessionId then get it
        uint256 sessionId = _sessionIdCounter.current();
        _sessionIdCounter.increment();

        // 3. Add new UploadSession with NO confirmed
        sessions[sessionId] = UploadSession({
            id: sessionId,
            user: msg.sender,
            proof: "",  // Proof will be confirmed later in `confirm()`
            confirmed: false
        });

        // 4. Mark `docId` is submitted
        docSubmits[docId] = true;

        // 5. Emit UploadData event
        emit UploadData(docId, sessionId);

        // 6. Return sessionId
        return sessionId;
    }

    function confirm(
        string memory docId,
        string memory contentHash,
        string memory proof,
        uint256 sessionId,
        uint256 riskScore
    ) public returns (uint256) {
        // 1. Verify session
        UploadSession storage session = sessions[sessionId];
        require(session.user == msg.sender, "Invalid session owner");
        require(docSubmits[docId], "Session is ended");
        require(!session.confirmed, "Doc already been submitted");

        // 2. TODO: Verify proof, we can skip this step
        // require(keccak256(bytes(proof)) == keccak256(bytes("success")), "Invalid proof");

        // 3. Update doc content
        docs[docId] = DataDoc({
            id: docId,
            hashContent: contentHash
        });

        // 4. Mint NFT and link `nftDocs` with `docId`
        uint256 nftId = geneNFT.safeMint(session.user);
        nftDocs[nftId] = docId;
        emit MintedGNFT(docId, sessionId, nftId);

        // 5. Reward PCSP token based on risk stroke
        pcspToken.reward(session.user, riskScore);
        emit RewardedPCSP(docId, sessionId, riskScore);

        // 6. Update proof and close session
        session.proof = proof;
        session.confirmed = true;

        return nftId;
    }

    function getSession(uint256 sessionId) public view returns (UploadSession memory) {
        return sessions[sessionId];
    }

    function getDoc(string memory docId) public view returns (DataDoc memory) {
        return docs[docId];
    }

    // Access Control: Start ===========================================================================================

    // Access Control: NFT ID -> user address -> access granted (true/false)
    mapping(uint256 => mapping(address => bool)) accessControl;

    // Events
    event GrantedAccess(uint256 nftId, address grantedTo);
    event RevokedAccess(uint256 nftId, address revokedFrom);

    // Modifier check permission to access GeneNFT
    modifier onlyNftOwner(uint256 nftId) {
        require(geneNFT.ownerOf(nftId) == msg.sender, "You are not the owner of this NFT.");
        _;
    }

    // Grant permission to an user
    function grantAccess(uint256 nftId, address user) public onlyNftOwner(nftId) {
        // 1. Update access = true
        accessControl[nftId][user] = true;
        // 2. Emit event
        emit GrantedAccess(nftId, user);
    }

    // Revoke permission from an user
    function revokeAccess(uint256 nftId, address user) public onlyNftOwner(nftId) {
        // 1. Update access = false
        accessControl[nftId][user] = false;
        // 2. Emit event
        emit RevokedAccess(nftId, user);
    }

    // Request access DataDoc
    function requestAccess(uint256 nftId) public view returns (DataDoc memory) {
        // 1. Validate permission
        require(
            geneNFT.ownerOf(nftId) == msg.sender || accessControl[nftId][msg.sender],
            "Access denied."
        );

        // 2. Get doc and return
        string memory docId = nftDocs[nftId];
        return docs[docId];
    }

    // Access Control: End ---------------------------------------------------------------------------------------------

}
