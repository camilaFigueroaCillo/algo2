class UnionFind:

    def __init__(self, elems):
        self.groups = {}
        for e in elems:
            self.groups[e] = e
    
    def find(self, e):
        if self.groups[e] == e:
            return e
        real_group = self.find(self.groups[e])
        self.groups[e] = real_group
        return real_group

    def union(self, e1, e2):
        new_group = self.find(e1)
        other = self.find(e2)
        self.groups[other] = new_group