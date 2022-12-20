package adapter

type OldClient struct {
}

func (r *OldClient) DoOldWork(b OldBusiness) {
	b.OldWork()
}

type NewClient struct {
}

func (n *NewClient) DoNewWork(b NewBusiness) {
	b.NewWork()
}
