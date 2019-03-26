package dynamic

type Pack struct {
	Weight int
	Value  int
	Obj    interface{}
}

type Knapsack struct {
	packs []Pack
	Memo  [][]int
}

func NewKnapsack(packs []Pack) *Knapsack {
	p := [][]int{}
	r := Knapsack{packs: packs, Memo: p}
	return &r
}

func (k *Knapsack) IndiceExists(i int, w int) bool {
	if len(k.Memo) < i {
		return false
	} else {
		if len(k.Memo[i]) < w {
			return false
		} else {
			return true
		}

	}
}

func (k *Knapsack) setIndices(i int, w int) {
	for len(k.Memo) <= i {
		k.Memo = append(k.Memo, []int{})
	}
	for len(k.Memo[i]) <= w {
		k.Memo[i] = append(k.Memo[i], -1)
	}

}

func (k *Knapsack) SolveProblem(i int, w int) int {
	if k.IndiceExists(i, w) && k.Memo[i][w] > -1 {
		return k.Memo[i][w]

	} else {
		if w == 0 || i == 0 {
			k.setIndices(i, w)
			k.Memo[i][w] = 0
			return k.Memo[i][w]
		}

		k.setIndices(i, w)
		i_pack := k.packs[i-1]

		if i_pack.Weight > w {
			k.Memo[i][w] = k.SolveProblem(i-1, w)
			return k.Memo[i][w]
		}

		solution_without_i := k.SolveProblem(i-1, w)
		solution_with_i := i_pack.Value + k.SolveProblem(i-1, w-i_pack.Weight)

		if solution_with_i > solution_without_i {
			k.Memo[i][w] = solution_with_i
		} else {
			k.Memo[i][w] = solution_without_i
		}
		return k.Memo[i][w]
	}

}
