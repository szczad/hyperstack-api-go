# Fix server address
.servers[0].url = "https://infrahub-api.nexgencloud.com/v1" |

# Fix schema definitions for certain components
.components.schemas."Instance Flavor Fields".properties.ram.type = "number" |
.components.schemas."Flavor Fields".properties.ram.type = "number" |
.components.schemas."Volume Fields".properties.size += { "type": "integer", "format": "int64" } |
.components.schemas."Image Fields".properties.size += { "format": "int64" } |

# Remove spaces in schema definitions to conform with the specs and fix references to these schemas
.components.schemas |= (to_entries | map({key: (.key | gsub(" |%20"; "")), value: .value}) | from_entries) | walk(if type == "object" then with_entries( if .key == "$ref" then .value |= gsub(" |%20"; "") end) else . end) |

# Remove duplicate tags
.tags |= (map(.key = .name | .value = .description) | from_entries | to_entries | map(.name = .key | .description = .value | del(.key, .value))) |

# no-op
.
