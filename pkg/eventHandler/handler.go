package eventHandler

func ProcessCommand(channelName, userID, operation, operationGroup, operationArgs string) string {
	return channelName + " " + userID + " " + operation + " " + operationGroup + " " + operationArgs
}
