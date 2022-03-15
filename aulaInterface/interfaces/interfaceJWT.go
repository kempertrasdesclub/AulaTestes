package interfaces

// InterfaceJWT (Português): Interface do algoritmo de JWT usado para validar sessão do usuário.
// tokenUID é usado como token de sessão e pode ser apagado para apagar uma sessão de usuário logado.
type InterfaceJWT interface {
	// NewAlgorithm (Português): Inicializa o algoritmo com a chave
	//   secretKey: chave única usada na criptografia do token
	//   err: objeto de erro
	NewAlgorithm(secretKey []byte) (err error)

	// BuildToken (Português):
	//   userUID: ID único do usuário
	//   tokenUID: ID único do token, usado para validar se o token ainda é válido no sistema
	//   audience: URL do site
	//   toke: token JWT
	//   err: objeto de erro
	BuildToken(userUID, tokenUID string, audience []string) (token []byte, err error)

	// Verify (Português): Verifica o token JWT
	//   token: token JWT
	//   tokenUID: ID único do token, usado para validar se o token ainda é válido no sistema
	//   userUID: ID único do usuário
	Verify(token []byte) (tokenUID, userUI string, err error)
}
