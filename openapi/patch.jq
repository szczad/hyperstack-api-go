# Fix schema definitions for certain components
.components.schemas."Instance Flavor Fields".properties.ram.type = "number" |
.components.schemas."Flavor Fields".properties.ram.type = "number" |
.components.schemas."Volume Fields".properties.size.type = "integer" |

# Remove spaces in schema definitions to conform with the specs and fix references to these schemas
.components.schemas |= (to_entries | map({key: (.key | gsub(" |%20"; "")), value: .value}) | from_entries) | walk(if type == "object" then with_entries( if .key == "$ref" then .value |= gsub(" |%20"; "") end) else . end) |

# Remove duplicate tags
.tags |= (map(.key = .name | .value = .description) | from_entries | to_entries | map(.name = .key | .description = .value | del(.key, .value))) |

# no-op
.
