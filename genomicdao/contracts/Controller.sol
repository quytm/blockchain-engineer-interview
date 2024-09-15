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
    event UploadData(string docId, uint256 sessionId);

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
    ) public {
        // 1. Verify session
        UploadSession storage session = sessions[sessionId];
        require(session.user == msg.sender, "Invalid session owner");
        require(docSubmits[docId], "Session is ended");
        require(!session.confirmed, "Doc already been submitted");
        // require(bytes(session.proof).length == 0, "Session is ended");

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

        // 5. Reward PCSP token based on risk stroke
        pcspToken.reward(session.user, riskScore);

        // 6. Update proof and close session
        session.proof = proof;
        session.confirmed = true;
    }

    function getSession(uint256 sessionId) public view returns (UploadSession memory) {
        return sessions[sessionId];
    }

    function getDoc(string memory docId) public view returns (DataDoc memory) {
        return docs[docId];
    }
}
