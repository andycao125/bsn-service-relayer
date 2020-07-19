package hub

// AddKey implements KeyManager
func (ic IritaHubChain) AddKey(name string, passphrase string) (addr string, mnemonic string, err error) {
	return ic.Client.Keys.Add(name, passphrase)
}

// DeleteKey implements KeyManager
func (ic IritaHubChain) DeleteKey(name string) error {
	return ic.Client.Keys.Delete(name)
}

// GetKey implements KeyManager
func (ic IritaHubChain) GetKey(name string) (addr string, err error) {
	return ic.Client.Keys.Show(name)
}

// ImportKey implements KeyManager
func (ic IritaHubChain) ImportKey(name string, passphrase string, keystore string) (addr string, err error) {
	return ic.Client.Keys.Import(name, passphrase, keystore)
}

// ExportKey implements KeyManager
func (ic IritaHubChain) ExportKey(name string, passphrase string, encryptPassphrase string) (mnemonic string, err error) {
	return ic.Client.Keys.Export(name, passphrase, encryptPassphrase)
}
