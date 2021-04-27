# account-microservice

## Comment lancer et compiler le microservice

Pour compiler, aller à la racine du projet et utiliser la commande 
```powershell
go build
```
Cette commande produit un exécutable qu'il suffit de lancer pour que le microservice soit actif.

## Comment accéder au microservice

Ce microservice se lance sur localhost:8000 par défaut. Pour en changer la configuration, modifiez le fichier main.go à la ligne 29 :
```go
err := http.ListenAndServe(":<port>", accountService.MakeHTTPHandler(service, logger))
```

## Comment paramétrer l'accès à la base de données

Pour paramétrer l'accès à la base de données il suffit de modifier la structure info présente dans le [main](https://github.com/PP-Groupe-6/account-microservice/blob/master/main.go) :
```go
	info := accountService.DbConnexionInfo{
		DbUrl:    "postgre://",
		DbPort:   "5432",
		DbName:   "prix_banque_test",
		Username: "dev",
		Password: "dev",
	}
```

Pour tester le microservice nous conseillons l'outil [Postman](https://www.postman.com) et [la collection fournie avec le microservice](https://github.com/PP-Groupe-6/account-microservice/blob/master/Account.postman_collection.json).

La liste des Url est la suivante :
| URL                     | Méthode           | Param (JSON dans le body) | Retour               |
| ----------------------- |:-----------------:| :------------------------:| :-------------------:|
| localhost:8000/amount/  | GET               | {"ClientID": "\<ID\>"}      |{"amount": \<amount\>}  |
| localhost:8000/users/   | GET               | {"ClientID": "\<ID\>"}      | {"fullName": "\<fullname\>","mailAdress": "\<mail\>","phoneNumber": "\<number\>"}|
| localhost:8000/users/   | POST              | {"ClientID":"\<Firebase ID\>","FullName":"\<fullname\>","PhoneNumber":"\<number\>","MailAdress":"\<mail\>"}| {"fullName": "\<fullname\>","mailAdress": "\<mail\>","phoneNumber": "\<number\>"}| {"accepted": "\<bool\>"}|
