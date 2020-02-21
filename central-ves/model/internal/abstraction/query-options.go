package abstraction

type QueryOptionOrder struct {
	Order   string
	Reorder bool
}

func (o QueryOptionOrder) implementsChainInfoQuery() ChainInfoQueryOption { return o }
func (o QueryOptionOrder) implementsObjectQuery() ObjectQueryOption       { return o }
func (o QueryOptionOrder) implementsUserQuery() UserQueryOption           { return o }

type QueryOptionPage struct {
	Page     int
	PageSize int
}

func (o QueryOptionPage) implementsChainInfoQuery() ChainInfoQueryOption { return o }
func (o QueryOptionPage) implementsObjectQuery() ObjectQueryOption       { return o }
func (o QueryOptionPage) implementsUserQuery() UserQueryOption           { return o }

type QueryOptionBeforeID struct {
	ID int
}

func (o QueryOptionBeforeID) implementsChainInfoQuery() ChainInfoQueryOption { return o }
func (o QueryOptionBeforeID) implementsObjectQuery() ObjectQueryOption       { return o }
func (o QueryOptionBeforeID) implementsUserQuery() UserQueryOption           { return o }

type QueryOptionPreload struct{}

func (o QueryOptionPreload) implementsChainInfoQuery() ChainInfoQueryOption { return o }
func (o QueryOptionPreload) implementsObjectQuery() ObjectQueryOption       { return o }
func (o QueryOptionPreload) implementsUserQuery() UserQueryOption           { return o }
