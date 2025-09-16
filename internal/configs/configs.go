package configs

import "fmt"

type HttpConfig struct {
	Port uint16
}

func (h *HttpConfig) GetStringPort() string {
	return fmt.Sprintf(":%d", h.Port)

}
