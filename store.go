package docgen

import (
	"fmt"
	"log"

	"golang.org/x/net/context"

	firebase "firebase.google.com/go"
	_ "firebase.google.com/go/auth"

	"google.golang.org/api/option"
)

func InitFireStore() (*firebase.App, error) {
	opt := option.WithCredentialsFile("./serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return app, nil
}

func GetDataByID(app *firebase.App, id string) (map[string]interface{}, error) {
	ctx := context.Background()
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	dsnap, err := client.Collection("docs").Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}
	m := dsnap.Data()
	fmt.Printf("Document data: %#v\n", m)
	return m, nil
}
