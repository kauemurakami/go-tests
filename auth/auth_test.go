package auth

import (
	"fmt"
	"testing"
)

func TestAuth(t *testing.T) {
	uPass := Auth(user{"kauê", "123"})
	if uPass {
		t.Log("Usuário logado")
		fmt.Println("Usuário logado")
	} else {
		t.Errorf("Erro ao logar")
	}

}
