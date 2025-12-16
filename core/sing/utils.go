package sing

import (
	"fmt"
	"strconv"

	"github.com/InazumaV/V2bX/conf"
	"github.com/sagernet/sing-box/option"
)

func processFallback(c *conf.Options, fallbackForALPN map[string]*option.ServerOptions) error {
	for k, v := range c.SingOptions.FallBackConfigs.FallBackForALPN {
		fallbackPort, err := strconv.Atoi(v.ServerPort)
		if err != nil {
			return fmt.Errorf("unable to parse fallbackForALPN server port error: %s", err)
		}
		if fallbackPort <= 0 || fallbackPort > 65535 {
			return fmt.Errorf("invalid fallbackForALPN port: %d", fallbackPort)
		}
		if v.Server == "" {
			return fmt.Errorf("fallbackForALPN server is empty")
		}
		fallbackForALPN[k] = &option.ServerOptions{
			Server:     v.Server,
			ServerPort: uint16(fallbackPort),
		}
	}
	return nil
}
