package v1

//go:generate go-enum -f=$GOFILE --marshal --names

// ListResponseResultFirstMileTypeFirstMileType
// ENUM(DropOff=1, Pickup, DIRECTION_RISE, DIRECTION_FALL)
type ListResponseResultFirstMileTypeFirstMileType int32

// ListResponseResultStatus
// ENUM(new=1, created, disabled, blocked, disabled_due_to_limit, error)
type ListResponseResultStatus int32
