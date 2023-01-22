package errors

type HashMismatch string

func (chm HashMismatch) Error() string {
	return "hashMismatch: " + string(chm)
}

type SignatureNotFound string

func (snf SignatureNotFound) Error() string {
	return "signatureNotFound: " + string(snf)
}

type InvalidIP string

func (ipi InvalidIP) Error() string {
	return "invalidIP: " + string(ipi)
}

type PublicKeyNotFound string

func (pnf PublicKeyNotFound) Error() string {
	return "publicKeyNotFound: " + string(pnf)
}

type InvalidSave string

func (is InvalidSave) Error() string {
	return "invalidSave: " + string(is)
}
