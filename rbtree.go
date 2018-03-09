package main

type ngx_rbtree_key_t uint64

type ngx_rbtree_insert_pt func(root, node, sentinel *ngx_rbtree_node_t)

type ngx_rbtree_node_t struct {
	key    ngx_rbtree_key_t
	left   *ngx_rbtree_node_t
	right  *ngx_rbtree_node_t
	parent *ngx_rbtree_node_t
	color  uint8
	data   interface{}
}

type ngx_rbtree_t struct {
	root     *ngx_rbtree_node_t
	sentinel *ngx_rbtree_node_t
	insert   ngx_rbtree_insert_pt
}

func ngx_rbtree_sentinel_init(node *ngx_rbtree_node_t) {
	ngx_rbt_black(node)
}

func ngx_rbt_red(node *ngx_rbtree_node_t) {
	node.color = 1
}

func ngx_rbt_black(node *ngx_rbtree_node_t) {
	node.color = 0
}

func ngx_rbt_is_red(node *ngx_rbtree_node_t) bool {
	return node.color == 1
}

func ngx_rbt_is_black(node *ngx_rbtree_node_t) bool {
	return node.color == 0
}

func ngx_rbt_copy_color(n1, n2 *ngx_rbtree_node_t) {
	n1.color = n2.color
}

func ngx_rbtree_left_rotate(root **ngx_rbtree_node_t, sentinel, node *ngx_rbtree_node_t) {

	temp := node.right
	node.right = temp.left

	if temp.left != sentinel {
		temp.left.parent = node
	}

	temp.parent = node.parent

	if node == *root {
		*root = temp
	} else if node == node.parent.left {
		node.parent.left = temp
	} else {
		node.parent.right = temp
	}

	temp.left = node
	node.parent = temp
}

func ngx_rbtree_right_rotate(root **ngx_rbtree_node_t, sentinel, node *ngx_rbtree_node_t) {

	temp := node.left
	node.left = temp.right

	if temp.right != sentinel {
		temp.right.parent = node
	}

	temp.parent = node.parent

	if node == *root {
		*root = temp
	} else if node == node.parent.right {
		node.parent.right = temp
	} else {
		node.parent.left = temp
	}

	temp.right = node
	node.parent = temp
}

func ngx_rbtree_min(node, sentinel *ngx_rbtree_node_t) *ngx_rbtree_node_t {
	for node.left != sentinel {
		node = node.left
	}
	return node
}

func ngx_rbtree_init(tree *ngx_rbtree_t, s *ngx_rbtree_node_t, i ngx_rbtree_insert_pt) {
	ngx_rbtree_sentinel_init(s)
	tree.root = s
	tree.sentinel = s
	tree.insert = i
}

func ngx_rbtree_insert(tree *ngx_rbtree_t, node *ngx_rbtree_node_t) {

	var (
		root     **ngx_rbtree_node_t
		temp     *ngx_rbtree_node_t
		sentinel *ngx_rbtree_node_t
	)

	/* a binary tree insert */

	root = &tree.root
	sentinel = tree.sentinel

	if *root == sentinel {
		node.parent = nil
		node.left = sentinel
		node.right = sentinel
		ngx_rbt_black(node)
		*root = node
		return
	}

	tree.insert(*root, node, sentinel)

	/* re-balance tree */

	for node != *root && ngx_rbt_is_red(node.parent) {

		if node.parent == node.parent.parent.left {
			temp = node.parent.parent.right

			if ngx_rbt_is_red(temp) {
				ngx_rbt_black(node.parent)
				ngx_rbt_black(temp)
				ngx_rbt_red(node.parent.parent)
				node = node.parent.parent

			} else {
				if node == node.parent.right {
					node = node.parent
					ngx_rbtree_left_rotate(root, sentinel, node)
				}

				ngx_rbt_black(node.parent)
				ngx_rbt_red(node.parent.parent)
				ngx_rbtree_right_rotate(root, sentinel, node.parent.parent)
			}

		} else {
			temp = node.parent.parent.left

			if ngx_rbt_is_red(temp) {
				ngx_rbt_black(node.parent)
				ngx_rbt_black(temp)
				ngx_rbt_red(node.parent.parent)
				node = node.parent.parent

			} else {
				if node == node.parent.left {
					node = node.parent
					ngx_rbtree_right_rotate(root, sentinel, node)
				}

				ngx_rbt_black(node.parent)
				ngx_rbt_red(node.parent.parent)
				ngx_rbtree_left_rotate(root, sentinel, node.parent.parent)
			}
		}
	}

	ngx_rbt_black(*root)
}

func ngx_rbtree_insert_value(temp, node, sentinel *ngx_rbtree_node_t) {
	var p **ngx_rbtree_node_t

	for {

		if node.key < temp.key {
			p = &temp.left
		} else {
			p = &temp.right
		}

		if *p == sentinel {
			break
		}

		temp = *p
	}

	*p = node
	node.parent = temp
	node.left = sentinel
	node.right = sentinel
	ngx_rbt_red(node)
}

func ngx_rbtree_delete(tree *ngx_rbtree_t, node *ngx_rbtree_node_t) {
	var (
		red      bool
		root     **ngx_rbtree_node_t
		sentinel *ngx_rbtree_node_t
		subst    *ngx_rbtree_node_t
		temp     *ngx_rbtree_node_t
		w        *ngx_rbtree_node_t
	)

	/* a binary tree delete */

	root = &tree.root
	sentinel = tree.sentinel

	if node.left == sentinel {
		temp = node.right
		subst = node
	} else if node.right == sentinel {
		temp = node.left
		subst = node
	} else {
		subst = ngx_rbtree_min(node.right, sentinel)
		if subst.left != sentinel {
			temp = subst.left
		} else {
			temp = subst.right
		}
	}

	if subst == *root {
		*root = temp
		ngx_rbt_black(temp)

		/* DEBUG stuff */
		node.left = nil
		node.right = nil
		node.parent = nil
		node.key = 0

		return
	}

	red = ngx_rbt_is_red(subst)

	if subst == subst.parent.left {
		subst.parent.left = temp
	} else {
		subst.parent.right = temp
	}

	if subst == node {
		temp.parent = subst.parent
	} else {
		if subst.parent == node {
			temp.parent = subst
		} else {
			temp.parent = subst.parent
		}

		subst.left = node.left
		subst.right = node.right
		subst.parent = node.parent
		ngx_rbt_copy_color(subst, node)

		if node == *root {
			*root = subst
		} else {
			if node == node.parent.left {
				node.parent.left = subst
			} else {
				node.parent.right = subst
			}
		}

		if subst.left != sentinel {
			subst.left.parent = subst
		}

		if subst.right != sentinel {
			subst.right.parent = subst
		}
	}

	/* DEBUG stuff */
	node.left = nil
	node.right = nil
	node.parent = nil
	node.key = 0

	if red {
		return
	}

	/* a delete fixup */

	for temp != *root && ngx_rbt_is_black(temp) {

		if temp == temp.parent.left {
			w = temp.parent.right

			if ngx_rbt_is_red(w) {
				ngx_rbt_black(w)
				ngx_rbt_red(temp.parent)
				ngx_rbtree_left_rotate(root, sentinel, temp.parent)
				w = temp.parent.right
			}

			if ngx_rbt_is_black(w.left) && ngx_rbt_is_black(w.right) {
				ngx_rbt_red(w)
				temp = temp.parent
			} else {
				if ngx_rbt_is_black(w.right) {
					ngx_rbt_black(w.left)
					ngx_rbt_red(w)
					ngx_rbtree_right_rotate(root, sentinel, w)
					w = temp.parent.right
				}
				ngx_rbt_copy_color(w, temp.parent)
				ngx_rbt_black(temp.parent)
				ngx_rbt_black(w.right)
				ngx_rbtree_left_rotate(root, sentinel, temp.parent)
				temp = *root
			}

		} else {
			w = temp.parent.left

			if ngx_rbt_is_red(w) {
				ngx_rbt_black(w)
				ngx_rbt_red(temp.parent)
				ngx_rbtree_right_rotate(root, sentinel, temp.parent)
				w = temp.parent.left
			}

			if ngx_rbt_is_black(w.left) && ngx_rbt_is_black(w.right) {
				ngx_rbt_red(w)
				temp = temp.parent
			} else {
				if ngx_rbt_is_black(w.left) {
					ngx_rbt_black(w.right)
					ngx_rbt_red(w)
					ngx_rbtree_left_rotate(root, sentinel, w)
					w = temp.parent.left
				}

				ngx_rbt_copy_color(w, temp.parent)
				ngx_rbt_black(temp.parent)
				ngx_rbt_black(w.left)
				ngx_rbtree_right_rotate(root, sentinel, temp.parent)
				temp = *root
			}
		}
	}

	ngx_rbt_black(temp)
}

func travel_rbtree(root, sentinel *ngx_rbtree_node_t, f func(key ngx_rbtree_key_t, data interface{})) {
	if root.left != sentinel {
		travel_rbtree(root.left, sentinel, f)
	}
	f(root.key, root.data)
	if root.right != sentinel {
		travel_rbtree(root.right, sentinel, f)
	}
}

type RbTree struct {
	rbtree   *ngx_rbtree_t
	sentinel *ngx_rbtree_node_t
}

func NewRbTree() *RbTree {
	rbmap := &RbTree{
		rbtree:   &ngx_rbtree_t{},
		sentinel: &ngx_rbtree_node_t{},
	}
	ngx_rbtree_init(rbmap.rbtree, rbmap.sentinel, ngx_rbtree_insert_value)
	return rbmap
}

func (this *RbTree) Store(key ngx_rbtree_key_t, value interface{}) {
	ngx_rbtree_insert(this.rbtree, &ngx_rbtree_node_t{key: key, data: value})
}

func (this *RbTree) Load(key ngx_rbtree_key_t) (value interface{}, ok bool) {
	tmpnode := this.rbtree.root
	for tmpnode != this.sentinel {
		if key == tmpnode.key {
			return tmpnode.data, true
		}
		if key < tmpnode.key {
			tmpnode = tmpnode.left
		} else {
			tmpnode = tmpnode.right
		}
	}
	return nil, false
}

func (this *RbTree) Delete(key ngx_rbtree_key_t) bool {
	tmpnode := this.rbtree.root
	for tmpnode != this.sentinel {
		if key == tmpnode.key {
			ngx_rbtree_delete(this.rbtree, tmpnode)
			return true
		}
		if key < tmpnode.key {
			tmpnode = tmpnode.left
		} else {
			tmpnode = tmpnode.right
		}
	}
	return false
}

func (this *RbTree) Range(f func(key ngx_rbtree_key_t, value interface{})) {
	travel_rbtree(this.rbtree.root, this.sentinel, f)
}
