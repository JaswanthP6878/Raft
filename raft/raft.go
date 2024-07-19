package raft

type Raft struct {
	currentTerm int
	votedFor    int
	log         []int // log type is temporary set as int

	commitIndex int
	lastApplied int

	// Volatile state for leaders (reset after election)
	nextIndex  []int // for each server, index of the next log entry to send to that server
	matchIndex []int // for each server, index of highest log entry known to be replicated on server
}

// constructor
func new() {
	return Raft{
		currentTerm: 0,
		votedFor:    -1,
		log:         []int{},
		commitIndex: 0,
		lastApplied: 0,
	}
}

func (r *Raft) AppendEntriesRPC(args AppendEntriesArgs, result *AppendEntriesResult) error {
	return nil // implementation is pending
}

func (r *Raft) RequestVoteRPC(args RequestVoteArgs, result *RequestVoteReply) error {
	return nil // implementation is pending

}
