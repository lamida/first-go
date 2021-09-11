
class Tree:
    def __init__(self, value, left=None, right=None):
        self.value = value
        self.left = left
        self.right = right

    def __str__(self):
        return str(self.value)


def walk(root: Tree):
    if root is None:
        return
    print(root.value)
    walk(root.left)
    walk(root.right)

if __name__  == "__main__":
    t = Tree(1, Tree(2, Tree(4), Tree(5)), Tree(3))
    walk(t)