test:
	revel test -v -a booking
	revel test -v -a booking2
	revel test -v -a chat
	revel test -v -a facebook-oauth2
	revel test -v -a i18n
	revel test -v -a orm/gorm
	revel test -v -a swagger
	revel test -v -a twitter-oauth
	revel test -v -a upload
	revel test -v -a validation

