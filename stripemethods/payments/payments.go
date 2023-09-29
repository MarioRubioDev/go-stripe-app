package payments

import "go-stripe-app/stripemethods"

var stripeKey string = stripemethods.DotEnvVariable("STRIPE_KEY")
