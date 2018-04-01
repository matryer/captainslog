package captainslog

import (
	"context"
	"testing"
)

func TestLog(t *testing.T) {
	ctx := context.Background()
	log(ctx, "Mr Worf did another shit on the bridge. We got Mr Data to clean it up because he doesn't mind the smell.")
	log(ctx, "It was nice to have Georgie on the bridge again while Mr Data was in the shop for a deep clean.")
}
