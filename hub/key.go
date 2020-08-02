package hub

// AddKey implements KeyManager
func (ic IritaHubChain) AddKey(name string, passphrase string) (addr string, mnemonic string, err error) {
	return ic.Client.Key.Add(name, passphrase)
}

// DeleteKey implements KeyManager
func (ic IritaHubChain) DeleteKey(name string, passphrase string) error {
	return ic.Client.Key.Delete(name, passphrase)
}

// ShowKey implements KeyManager
func (ic IritaHubChain) ShowKey(name string, passphrase string) (addr string, err error) {
	return ic.Client.Key.Show(name, passphrase)
}

// ImportKey implements KeyManager
func (ic IritaHubChain) ImportKey(name string, passphrase string, keyArmor string) (addr string, err error) {
	return ic.Client.Key.Import(name, passphrase, keyArmor)
}

// ExportKey implements KeyManager
func (ic IritaHubChain) ExportKey(name string, passphrase string) (keyArmor string, err error) {
	return ic.Client.Key.Export(name, passphrase)
}

// RecoverKey implements KeyManager
func (ic IritaHubChain) RecoverKey(name string, passphrase string, mnemonic string) (addr string, err error) {
	return ic.Client.Key.Recover(name, passphrase, mnemonic)
}
