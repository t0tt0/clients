package abstraction

type QueryOptionOrder struct {
	Order   string
	Reorder bool
}

func (o QueryOptionOrder) implementsObjectQuery() ObjectQueryOption                 { return o }
func (o QueryOptionOrder) implementsSessionQuery() SessionQueryOption               { return o }
func (o QueryOptionOrder) implementsSessionAccountQuery() SessionAccountQueryOption { return o }
func (o QueryOptionOrder) implementsTransactionQuery() TransactionQueryOption       { return o }

type QueryOptionPage struct {
	Page     int
	PageSize int
}

func (o QueryOptionPage) implementsObjectQuery() ObjectQueryOption                 { return o }
func (o QueryOptionPage) implementsSessionQuery() SessionQueryOption               { return o }
func (o QueryOptionPage) implementsSessionAccountQuery() SessionAccountQueryOption { return o }
func (o QueryOptionPage) implementsTransactionQuery() TransactionQueryOption       { return o }

type QueryOptionBeforeID struct {
	ID int
}

func (o QueryOptionBeforeID) implementsObjectQuery() ObjectQueryOption                 { return o }
func (o QueryOptionBeforeID) implementsSessionQuery() SessionQueryOption               { return o }
func (o QueryOptionBeforeID) implementsSessionAccountQuery() SessionAccountQueryOption { return o }
func (o QueryOptionBeforeID) implementsTransactionQuery() TransactionQueryOption       { return o }

type QueryOptionPreload struct{}

func (o QueryOptionPreload) implementsObjectQuery() ObjectQueryOption                 { return o }
func (o QueryOptionPreload) implementsSessionQuery() SessionQueryOption               { return o }
func (o QueryOptionPreload) implementsSessionAccountQuery() SessionAccountQueryOption { return o }
func (o QueryOptionPreload) implementsTransactionQuery() TransactionQueryOption       { return o }
