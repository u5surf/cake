package capv

// RequiredCommands checks the PATH for required commands
func (mc *MgmtCluster) RequiredCommands() []string {
	if !mc.Addons.Solidfire.Enable {
		RequiredCommands.Remove(string(tridentctl))
	}
	if !mc.Addons.Observability.Enable {
		RequiredCommands.Remove(string(helm))
	}

	return RequiredCommands.Exist()
}
