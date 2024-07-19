package raft

type AppendEntriesArgs struct {
	term         int
	leaderId     int
	prevLogIndex int
	prevLogTerm  int

	entries      []int
	leaderCommit int
}

type AppendEntriesResult struct {
	term    int
	success bool
}

type RequestVoteArgs struct {
	term         int
	candidateId  int
	lastLogIndex int
	lastLogTerm  int
}

type RequestVoteReply struct {
	term        int
	voteGranted bool
}
