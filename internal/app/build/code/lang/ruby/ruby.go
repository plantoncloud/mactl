package ruby

func GetEnvVars() map[string]string {
	return map[string]string{
		"LDFLAGS":         "-L/usr/local/opt/ruby/lib",
		"CPPFLAGS":        "-I/usr/local/opt/ruby/include",
		"PKG_CONFIG_PATH": "/usr/local/opt/ruby/lib/pkgconfig",
	}
}
