package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

type Node struct {
	key interface{}
	value interface{}
	Parent *Node
	Left *Node
	Right *Node
}

type SplayTree interface {
	SetRoot (n*Node)
	GetRoot ()*Node
	Ord (key1,key2 interface{}) int  //排序
}

//伸展树
type MySplayTree struct {
	root *Node

}

func (ST *MySplayTree) SetRoot (n*Node) {
	ST.root = n

}

func (ST *MySplayTree) GetRoot ()*Node {
	return ST.root

}

func (ST *MySplayTree)Ord(key1,key2 interface{})int  {
	if key1.(int) < key2.(int){
		return -1
	}else if key1.(int) == key2.(int){
		return 0
	}else {
		return 1
	}
}

func Search(ST SplayTree,key interface{}) *Node {
	return SearchNode(ST,key,ST.GetRoot())
	//实现二叉树查找
}

func SearchNode(ST SplayTree,key interface{},n*Node) *Node  {
	if n == nil{
		return nil

	}else {
		switch ST.Ord(key,n.key) {
		case 0:
			return n //相等返回n
		case -1:
			return SearchNode(ST,key,n.Left) //返回左子树
		case 1:
			return SearchNode(ST,key,n.Right)
		}
		return nil
	}
}

func Find(ST SplayTree,key interface{}) interface{} {
	return FindNode(ST,key,ST.GetRoot())
	//实现二叉树查找
}

func FindNode(ST SplayTree,key interface{},n*Node) interface{}  {
	if n == nil{
		return nil

	}else {
		switch ST.Ord(key,n.key) {
		case 0:
			return n.value //相等返回n
		case -1:
			return FindNode(ST,key,n.Left) //返回左子树
		case 1:
			return FindNode(ST,key,n.Right)
		}
		return nil
	}
}

//对Insert函数 封装
func Insert(ST SplayTree,key interface{},value interface{}) error{
	if Search(ST,key)!=nil{
		return errors.New("已经存在")
	}
	n := InsertNode(ST,key,value,ST.GetRoot()) //调用插入
	fmt.Println(n)
	return nil
}


//插入数据
func InsertNode(ST SplayTree,key interface{},value interface{},n *Node) *Node {
	if n == nil{
		_n := new(Node)
		_n.key = key
		_n.value=value
		ST.SetRoot(_n)
		return ST.GetRoot()
	}
	switch ST.Ord(key,n.key) {
	case 0:
		return nil //数据已经存在
	case -1:
		if n.Left==nil{
			n.Left 	= new(Node)
			n.Left.key = key
			n.Left.value=value
			n.Left.Parent=n
			return n.Left
		}else {
			return InsertNode(ST,key,value,n.Left)  //插入数据
		}
	case 1:
		if n.Right==nil{
			n.Right 	= new(Node)
			n.Right.key = key
			n.Right.value=value
			n.Right.Parent=n
			return n.Right

		}else {
			return InsertNode(ST,key,value,n.Right)  //插入数据
		}

	}
	return nil   //插入完成后旋转
}

//删除
func Delete(ST SplayTree,key interface{})error  {
	n:= Search(ST,key)
	if n==nil{
		return errors.New("请指定正确的删除对象")
	}else {
		//n相当于找到了位置
		p := n.Parent //保存父节点
		if n.Left != nil{
			iop := InOrderPredecessor(n.Left) //需按照左边最大
			Swap(n,iop) //交换节点
			Remove(ST,iop) //删除节点

		}else if n.Right!=nil{
			ios := InOrderSucccessor(n.Right)
			Swap(n,ios)
			Remove(ST,ios)

		}else {
			Remove(ST,n)

		}
		if p!=nil{
			//如果p不等于nil  需要伸展
		}
		return nil
	}
}

//删除节点
func Remove(ST SplayTree, n *Node)  {
	var isRoot bool
	var isLeft bool
	isRoot = (n==ST.GetRoot())  //判断是否根节点
	if isRoot!=true{
		isLeft=(n==n.Parent.Left)  //判断是否左节点
	}

	if isRoot != true{
		if isLeft ==true{
			if n.Left != nil{
				n.Parent.Left = n.Left
				n.Left.Parent = n.Parent

			}else if n.Right!=nil{
				n.Parent.Left = n.Right
				n.Right.Parent = n.Parent

			}else {
				n.Parent.Left = nil //叶子节点 左右都为空

			}
		}else {
			if n.Left != nil{
				n.Parent.Right = n.Left
				n.Left.Parent = n.Parent

			}else if n.Right!=nil{
				n.Parent.Right = n.Right
				n.Right.Parent = n.Parent

			}else {
				n.Parent.Right = nil

			}

		}
	}
	n = nil  //实现删除
}

func Swap(n1,n2 *Node)  {
	n1.key,n2.key=n2.key,n1.key
	n1.value,n2.value=n2.value,n1.value
}

//取得 极大极小值  用递归来实现
func InOrderPredecessor(n*Node)*Node{
	if n.Right==nil{
		return n
	}else{
		return InOrderPredecessor(n.Right)
	}
}
//取得最小
func InOrderSucccessor(n*Node)*Node{
	if n.Left==nil{
		return n
	}else{
		return InOrderSucccessor(n.Left)
	}
}
//伸展  左旋 右旋 实现一堆代码
func Splay(ST SplayTree,n*Node){
	for n!=ST.GetRoot(){
		if n.Parent==ST.GetRoot() && n.Parent.Left==n{
			ZigL(ST ,n)//根节点
		}else if n.Parent==ST.GetRoot()&&n.Parent.Right==n{
			ZigR(ST ,n)//根节点
		}else if n.Parent.Left==n &&n.Parent.Parent.Left==n.Parent{
			ZigZigL(ST,n)
		}else if n.Parent.Right==n &&n.Parent.Parent.Right==n.Parent{
			ZigZigR(ST,n)
		}else if n.Parent.Right==n &&n.Parent.Parent.Left==n.Parent{
			ZigZigLR(ST,n)
		}else{
			ZigZigRL(ST,n)
		}
	}


}

func ZigL(ST SplayTree,n *Node)  {
	n.Parent.Left = n.Right  //存储右边的数据  保存到根节点   看图    O2
	//                                                           o1
	if n.Right != nil{
		n.Right.Parent = n.Parent
	}
	n.Parent.Parent = n   //这里做一个n的备份
	n.Right = n.Parent    //n = 2  需要把4变成n的左节点
	n.Parent = nil
	ST.SetRoot(n)


	//左旋是右旋的镜像
}

func ZigR(ST SplayTree,n *Node)  {


	n.Parent.Right = n.Left  //存储左边的数据  保存到根节点   看图    O2
	//                                                           o1
	if n.Left != nil{
		n.Left.Parent = n.Parent
	}
	n.Parent.Parent = n   //这里做一个n的备份
	n.Left = n.Parent    //n = 2  需要把4变成n的左节点
	n.Parent = nil
	ST.SetRoot(n)   //右旋示意图
	//   O4                  O2
	// O2  O6              O1    O4
	//O1 O3 O5 O7               O3  O6
	//                             O5 O7
}

func ZigZigL(ST SplayTree,n *Node)  {
	gg := n.Parent.Parent.Parent  //访问太爷爷
	var isRoot bool
	var isLeft bool
	if gg == nil{
		//太爷爷节点不存在
		isRoot = true
	}else {
		isRoot = false
		isLeft= (gg.Left == n.Parent.Parent)
	}

	n.Parent.Parent.Left = n.Parent.Right   //备份left
	if n.Parent.Right != nil{
		n.Parent.Right.Parent = n.Parent.Parent
	}
	n.Parent.Left= n.Right
	if n.Right != nil{
		n.Right.Parent = n.Parent
	}
	n.Parent.Right = n.Parent.Parent
	n.Parent.Parent.Parent = n.Parent
	n.Right = n.Parent
	n.Parent.Parent = n
	n.Parent = gg

	//判断树
	if isRoot == true{
		ST.SetRoot(n)  //如果是根节点直接赋值
	}else if isLeft == true{
		gg.Left = n
	}else {
		gg.Right = n
	}

	//双 左旋
}

func ZigZigR(ST SplayTree,n *Node)  {
	//双 右旋
	gg := n.Parent.Parent.Parent  //访问太爷爷
	var isRoot bool
	var isLeft bool
	if gg == nil{
		//太爷爷节点不存在
		isRoot = true
	}else {
		isRoot = false
		isLeft= (gg.Left == n.Parent.Parent)
	}

	n.Parent.Parent.Right = n.Parent.Left   //备份left
	if n.Parent.Left != nil{
		n.Parent.Right.Parent = n.Parent.Parent
	}
	n.Parent.Right= n.Left
	if n.Left != nil{
		n.Left.Parent = n.Parent
	}
	n.Parent.Left = n.Parent.Parent
	n.Parent.Parent.Parent = n.Parent
	n.Left = n.Parent
	n.Parent.Parent = n
	n.Parent = gg

	//判断树
	if isRoot == true{
		ST.SetRoot(n)  //如果是根节点直接赋值
	}else if isLeft == true{
		gg.Left = n
	}else {
		gg.Right = n
	}

}

func ZigZigLR(ST SplayTree,n *Node)  {
	//先左旋 再右旋
	gg := n.Parent.Parent.Parent  //访问太爷爷
	var isRoot bool
	var isLeft bool
	if gg == nil{
		//太爷爷节点不存在
		isRoot = true
	}else {
		isRoot = false
		isLeft= (gg.Left == n.Parent.Parent)
	}

	n.Parent.Parent.Left = n.Parent.Right   //备份left
	if n.Right != nil{
		n.Right.Parent = n.Parent.Parent
	}
	n.Parent.Right= n.Left
	if n.Left != nil{
		n.Left.Parent = n.Parent
	}
	n.Left = n.Parent
	n.Right = n.Parent.Parent
	n.Parent.Parent.Parent = n
	n.Parent.Parent = n
	n.Parent = gg

	//判断树
	if isRoot == true{
		ST.SetRoot(n)  //如果是根节点直接赋值
	}else if isLeft == true{
		gg.Left = n
	}else {
		gg.Right = n
	}

}
func ZigZigRL(ST SplayTree,n *Node)  {
	//先右旋  再左旋
	gg := n.Parent.Parent.Parent  //访问太爷爷
	var isRoot bool
	var isLeft bool
	if gg == nil{
		//太爷爷节点不存在
		isRoot = true
	}else {
		isRoot = false
		isLeft= (gg.Left == n.Parent.Parent)
	}

	n.Parent.Parent.Right = n.Parent.Left   //备份left
	if n.Left != nil{
		n.Left.Parent = n.Parent.Parent
	}
	n.Parent.Left= n.Right
	if n.Right != nil{
		n.Right.Parent = n.Parent
	}
	n.Right = n.Parent
	n.Left = n.Parent.Parent
	n.Parent.Parent.Parent = n
	n.Parent.Parent = n
	n.Parent = gg

	//判断树
	if isRoot == true{
		ST.SetRoot(n)  //如果是根节点直接赋值
	}else if isLeft == true{
		gg.Left = n
	}else {
		gg.Right = n
	}
}
//Strings

func Print(ST SplayTree)  {
	PrintNode(ST.GetRoot(),0) //打印树枝
}

func PrintNode(n *Node,level int)  {
	if n == nil{
		return  //实现一个中序遍历
	}
	PrintNode(n.Left,level+1)
	fmt.Println(strings.Repeat("-",2*level),n.key,n.value) //打印数据

	PrintNode(n.Right,level+1)
}

func main()  {
	ST:=new(MySplayTree)
	for i:=0;i<36;i++{
		err:=Insert(ST,rand.Int()%100," ")
		if err!=nil{
			fmt.Println(err)
		}else{
			Print(ST)
		}
	}
	for i:=0;i<36;i++{
		err:=Delete(ST,i)
		if err!=nil{
			fmt.Println(err)
		}else{
			Print(ST)
		}
	}

}