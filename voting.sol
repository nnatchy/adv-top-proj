// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract VotingSystem {
    struct Vote {
        string studentId;        // ID of the voter
        address voterAddress;    // Wallet address of the voter
        string voteChoice;       // Vote choice
    }

    Vote[] public allVotes;                 // Array to store all votes
    bool public votingOpen = true;          // Voting status

    address public owner; // Contract owner

    event VoteCast(string studentId, address voterAddress, string voteChoice);
    event VotingEnded();

    constructor() {
        owner = msg.sender; // Set the deployer as the contract owner
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Only the owner can perform this action.");
        _;
    }

    modifier votingActive() {
        require(votingOpen, "Voting has ended.");
        _;
    }

    // Cast a vote
    function castVote(string memory studentId, address voterAddress, string memory voteChoice) 
        public votingActive 
    {
        require(voterAddress == msg.sender, "Voter address must match the caller's address.");
        // Record the vote
        Vote memory newVote = Vote(studentId, voterAddress, voteChoice);
        allVotes.push(newVote); // Add vote to the array for full result tracking
        emit VoteCast(studentId, voterAddress, voteChoice);
    }

    // End voting (only callable by owner)
    function endVoting() public onlyOwner {
        require(votingOpen, "Voting is already closed.");
        votingOpen = false;
        emit VotingEnded();
    }

    // Retrieve all votes
    function getAllVotes() public view returns (Vote[] memory) {
        return allVotes;
    }
}
