package mongo

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-oauth2/oauth2/v4/models"
	. "github.com/smartystreets/goconvey/convey"
)

// shut down the database, test should fail within a second
func TestClientStoreWithTimeout(t *testing.T) {

	fmt.Println("Replica set is", isReplicaSet)

	storeConfig := NewStoreConfig(1, 1)

	var store *ClientStore
	store = NewClientStore(NewConfig(url, dbName, username, password, service, isReplicaSet), storeConfig)

	client := &models.Client{
		ID:     "id",
		Secret: "secret",
		Domain: "domain",
		UserID: "user_id",
	}

	Convey("Set", t, func() {
		Convey("HappyPath", func() {
			_ = store.RemoveByID(client.ID)

			err := store.Create(client)

			So(err, ShouldBeNil)
		})

		// in case the server restart the client will already exist
		// so it should not return an err
		Convey("AlreadyExistingClient", func() {
			_ = store.RemoveByID(client.ID)

			_ = store.Create(client)
			err := store.Create(client)

			So(err, ShouldBeNil)
		})
	})

	Convey("GetByID", t, func() {
		Convey("HappyPath", func() {
			_ = store.RemoveByID(client.ID)
			_ = store.Create(client)

			got, err := store.GetByID(context.TODO(), client.ID)

			So(err, ShouldBeNil)
			So(got, ShouldResemble, client)
		})

		Convey("UnknownClient", func() {
			_, err := store.GetByID(context.TODO(), "unknown_client")

			So(err.Error(), ShouldEqual, "mongo: no documents in result")
		})
	})

	Convey("RemoveByID", t, func() {
		Convey("UnknownClient", func() {

			// In case the document does not exist, the returned err is nil
			err := store.RemoveByID("unknown_client")

			So(err, ShouldBeNil)
		})
	})
}

func TestClientStore(t *testing.T) {
	var store *ClientStore
	store = NewClientStore(NewConfig(url, dbName, username, password, service, isReplicaSet))

	client := &models.Client{
		ID:     "id",
		Secret: "secret",
		Domain: "domain",
		UserID: "user_id",
	}

	Convey("Set", t, func() {
		Convey("HappyPath", func() {
			_ = store.RemoveByID(client.ID)

			err := store.Create(client)

			So(err, ShouldBeNil)
		})

		// in case the server restart the client will already exist
		// so it should not return an err
		Convey("AlreadyExistingClient", func() {
			_ = store.RemoveByID(client.ID)

			_ = store.Create(client)
			err := store.Create(client)

			So(err, ShouldBeNil)
		})
	})

	Convey("GetByID", t, func() {
		Convey("HappyPath", func() {
			_ = store.RemoveByID(client.ID)
			_ = store.Create(client)

			got, err := store.GetByID(context.TODO(), client.ID)

			So(err, ShouldBeNil)
			So(got, ShouldResemble, client)
		})

		Convey("UnknownClient", func() {
			_, err := store.GetByID(context.TODO(), "unknown_client")

			So(err.Error(), ShouldEqual, "mongo: no documents in result")
		})
	})

	Convey("RemoveByID", t, func() {
		Convey("UnknownClient", func() {

			// In case the document does not exist, the returned err is nil
			err := store.RemoveByID("unknown_client")

			So(err, ShouldBeNil)
		})
	})
}
