package operationhandler

var (
	SuccessfullyAdded   = New("SUCCESSFULLY_ADDED", "the item successfully added")
	SuccessfullyUpdated = New("SUCCESSFULLY_UPDATED", "the item successfuly updated")
)

type OperationHandler struct {
	Message     string `json:"message"`
	Description string `json:"description"`
}

func New(message, description string) *OperationHandler {
	return &OperationHandler{message, description}
}
