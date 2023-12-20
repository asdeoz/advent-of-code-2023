package models

type Node struct {
	Id        string `json:"id"`
	Left      string `json:"left"`
	Right     string `json:"right"`
	LeftNode  *Node
	RightNode *Node
}
