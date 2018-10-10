package flowhandler

var stragery_one *BaseHandler

const (
	ID_HANDLER         = 1
	IP_HANDLER         = 2
	SLOTID_HANDLER     = 3
	WHILTElIST_HANDLER = 4
)

//whiteList---->id---->ip--->slotid  (no combine)
func stragery_one_method() *BaseHandler {
	stragery_one = new(BaseHandler)
	whiteListHandler := new(WhiteListHandler)
	muidOverHandler := new(MuidOverFrequcentHandler)
	idhandler := new(IdHandler)
	iphandler := new(IpHandler)
	slothandler := new(SlotIdHandler)

	stragery_one.SetNext(whiteListHandler)
	whiteListHandler.SetNext(muidOverHandler)
	muidOverHandler.SetNext(idhandler)
	idhandler.SetNext(iphandler)
	iphandler.SetNext(slothandler)

	return stragery_one
}

func init() {
	stragery_one_method()
}
