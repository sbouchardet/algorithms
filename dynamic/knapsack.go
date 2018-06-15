package dynamic

type Pack struct {
	Weight int
	Value  int
	Obj    interface{}
}

type Knapsack struct {
	packs []Pack
	memo  [][]int
}

func NewKnapsack(packs []Pack) *Knapsack {
	p := [][]int{}
	r := Knapsack{packs: packs, memo: p}
	return &r
}

func (k *Knapsack) IndiceExists(i int, w int) bool {
	if len(k.memo) < i {
		return false
	} else {
		if len(k.memo[i]) < w {
			return false
		} else {
			return true
		}

	}
}

func (k *Knapsack) setIndices(i int, w int) {
	for len(k.memo) <= i {
		k.memo = append(k.memo, []int{})
	}
	for len(k.memo[i]) <= w {
		k.memo[i] = append(k.memo[i], -1)
	}

}

func (k *Knapsack) SolveProblem(i int, w int) int {
	if k.IndiceExists(i, w) && k.memo[i][w] > -1 {
		return k.memo[i][w]

	} else {
		if w == 0 || i == 0 {
			k.setIndices(i, w)
			k.memo[i][w] = 0
			return k.memo[i][w]
		}

		k.setIndices(i, w)
		i_pack := k.packs[i-1]

		if i_pack.Weight > w {
			k.memo[i][w] = k.SolveProblem(i-1, w)
			return k.memo[i][w]
		}

		solution_without_i := k.SolveProblem(i-1, w)
		solution_with_i := i_pack.Value + k.SolveProblem(i-1, w-i_pack.Weight)

		if solution_with_i > solution_without_i {
			k.memo[i][w] = solution_with_i
		} else {
			k.memo[i][w] = solution_without_i
		}
		return k.memo[i][w]
	}

}
