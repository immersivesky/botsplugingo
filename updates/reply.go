package updates

import (
	"encoding/json"
)

type Reply []Block
type Block []Action

type Action struct {
	Event	string	`json:"event,omitempty"`
	Connect	string	`json:"connect,omitempty"`
	Message	Message	`json:"message,omitempty"`
}

type ReplyJSON string

func CreateReply() *Reply {
	return &Reply{}
}

func(reply *Reply) JSON() ReplyJSON {
	replyJSON, _ := json.Marshal(reply)
	return ReplyJSON(replyJSON)
}

// ----------
// ~ BLOCKS ~
// ----------

func(blocks *Reply) Block(index int) *Block {
	if len(*blocks) >= index {
		blocks := *blocks
		return &blocks[index -1]
	}
	return &Block{}
}

func(blocks *Reply) AddBlock(action Action) *Reply {
	*blocks = append(*blocks, []Action{action})
	return blocks
}

func(block *Block) EditBlock(newBlock Block) *Block {
	*block = newBlock
	return block
}

func(blocks *Reply) DeleteBlock(index int) *Reply {
	index = index -1

	pointer := *blocks
    pointer[index] = pointer[len(pointer) -1]
    *blocks = pointer[:len(pointer) -1]

	return blocks
}

// -----------
// ~ ACTIONS ~
// -----------

func(block *Block) Action(index int) *Action {
	if len(*block) >= index {
		actions := *block
		return &actions[index -1]
	}
	return &Action{}
}

func(actions *Block) AddAction(action Action) {
	*actions = append(*actions, action)
}

func(action *Action) EditAction(newAction Action) {
	*action = newAction
}

func(actions *Block) DeleteAction(index int) *Block {
	pointer := *actions
	*actions = append(pointer[:index -1], pointer[index:]...)
	return actions
}