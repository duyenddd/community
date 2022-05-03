// Package apps provides a clean way for Tidbyt to be able to get a list of all
// community apps.
package apps

// Code generated by tools/generator. DO NOT EDIT.

import (
	"errors"

	"tidbyt.dev/community/apps/abstractclock"
	"tidbyt.dev/community/apps/ambientweather"
	"tidbyt.dev/community/apps/analogclock"
	"tidbyt.dev/community/apps/analogtime"
	"tidbyt.dev/community/apps/arcadeclassics"
	"tidbyt.dev/community/apps/baywheels"
	"tidbyt.dev/community/apps/bgghotness"
	"tidbyt.dev/community/apps/biblevotd"
	"tidbyt.dev/community/apps/bigclock"
	"tidbyt.dev/community/apps/binaryclock"
	"tidbyt.dev/community/apps/charlestownferry"
	"tidbyt.dev/community/apps/clockbyhenry"
	"tidbyt.dev/community/apps/coingeckoprice"
	"tidbyt.dev/community/apps/countdownclock"
	"tidbyt.dev/community/apps/cryptotracker"
	"tidbyt.dev/community/apps/dailykanji"
	"tidbyt.dev/community/apps/dateprogress"
	"tidbyt.dev/community/apps/datetimeclock"
	"tidbyt.dev/community/apps/daynightmap"
	"tidbyt.dev/community/apps/destiny2stats"
	"tidbyt.dev/community/apps/digibyteprice"
	"tidbyt.dev/community/apps/digitalrain"
	"tidbyt.dev/community/apps/duolingo"
	"tidbyt.dev/community/apps/dvdlogo"
	"tidbyt.dev/community/apps/dwheadline"
	"tidbyt.dev/community/apps/espnnews"
	"tidbyt.dev/community/apps/finevent"
	"tidbyt.dev/community/apps/fishbyt"
	"tidbyt.dev/community/apps/flags"
	"tidbyt.dev/community/apps/fuzzyclock"
	"tidbyt.dev/community/apps/gapilotbuddy"
	"tidbyt.dev/community/apps/goodservice"
	"tidbyt.dev/community/apps/happyhour"
	"tidbyt.dev/community/apps/hurricanetracker"
	"tidbyt.dev/community/apps/hvvdepartures"
	"tidbyt.dev/community/apps/ifparank"
	"tidbyt.dev/community/apps/isitchristmas"
	"tidbyt.dev/community/apps/isstracker"
	"tidbyt.dev/community/apps/jokesjokeapi"
	"tidbyt.dev/community/apps/launchcountdown"
	"tidbyt.dev/community/apps/lirr"
	"tidbyt.dev/community/apps/manifest"
	"tidbyt.dev/community/apps/mbta"
	"tidbyt.dev/community/apps/mbtanewtrains"
	"tidbyt.dev/community/apps/mlbleaders"
	"tidbyt.dev/community/apps/mlbscores"
	"tidbyt.dev/community/apps/mnlightrail"
	"tidbyt.dev/community/apps/moretransit"
	"tidbyt.dev/community/apps/moviequotes"
	"tidbyt.dev/community/apps/natdex"
	"tidbyt.dev/community/apps/nationaltoday"
	"tidbyt.dev/community/apps/nbascores"
	"tidbyt.dev/community/apps/nearearthobjs"
	"tidbyt.dev/community/apps/netatmo"
	"tidbyt.dev/community/apps/nflscores"
	"tidbyt.dev/community/apps/nft"
	"tidbyt.dev/community/apps/nhlnextgame"
	"tidbyt.dev/community/apps/nhlscores"
	"tidbyt.dev/community/apps/nightscout"
	"tidbyt.dev/community/apps/nixelclock"
	"tidbyt.dev/community/apps/noaabuoy"
	"tidbyt.dev/community/apps/nyancat"
	"tidbyt.dev/community/apps/nycbus"
	"tidbyt.dev/community/apps/ohhighwaysigns"
	"tidbyt.dev/community/apps/pagerduty"
	"tidbyt.dev/community/apps/pathtrainschedule"
	"tidbyt.dev/community/apps/phaseofmoon"
	"tidbyt.dev/community/apps/pokedex"
	"tidbyt.dev/community/apps/powerball"
	"tidbyt.dev/community/apps/purpleair"
	"tidbyt.dev/community/apps/randomslackmoji"
	"tidbyt.dev/community/apps/redditimages"
	"tidbyt.dev/community/apps/redditrplace"
	"tidbyt.dev/community/apps/sbbtimetable"
	"tidbyt.dev/community/apps/sfnextmuni"
	"tidbyt.dev/community/apps/snyk"
	"tidbyt.dev/community/apps/sportsscores"
	"tidbyt.dev/community/apps/sportsstandings"
	"tidbyt.dev/community/apps/spotthestation"
	"tidbyt.dev/community/apps/steam"
	"tidbyt.dev/community/apps/stepcounter"
	"tidbyt.dev/community/apps/strava"
	"tidbyt.dev/community/apps/sunrisesunset"
	"tidbyt.dev/community/apps/surflive"
	"tidbyt.dev/community/apps/tempest"
	"tidbyt.dev/community/apps/theysaidso"
	"tidbyt.dev/community/apps/todoist"
	"tidbyt.dev/community/apps/traffic"
	"tidbyt.dev/community/apps/transsee"
	"tidbyt.dev/community/apps/twitterfollows"
	"tidbyt.dev/community/apps/unsplash"
	"tidbyt.dev/community/apps/usgsearthquakes"
	"tidbyt.dev/community/apps/usyieldcurve"
	"tidbyt.dev/community/apps/vergetaglines"
	"tidbyt.dev/community/apps/verticalmessage"
	"tidbyt.dev/community/apps/wantedposter"
	"tidbyt.dev/community/apps/warframecycles"
	"tidbyt.dev/community/apps/weathermap"
	"tidbyt.dev/community/apps/whosthatpokemon"
	"tidbyt.dev/community/apps/worldclock"
)

// GetManifests returns a list of all apps in the this repository. Add your applet
// below to include it in the Tidbyt Mobile app for all Tidbyt users.
func GetManifests() []manifest.Manifest {
	return []manifest.Manifest{
		abstractclock.New(),
		ambientweather.New(),
		analogclock.New(),
		analogtime.New(),
		arcadeclassics.New(),
		baywheels.New(),
		bgghotness.New(),
		biblevotd.New(),
		bigclock.New(),
		binaryclock.New(),
		charlestownferry.New(),
		clockbyhenry.New(),
		coingeckoprice.New(),
		countdownclock.New(),
		cryptotracker.New(),
		dailykanji.New(),
		dateprogress.New(),
		datetimeclock.New(),
		daynightmap.New(),
		destiny2stats.New(),
		digibyteprice.New(),
		digitalrain.New(),
		duolingo.New(),
		dvdlogo.New(),
		dwheadline.New(),
		espnnews.New(),
		finevent.New(),
		fishbyt.New(),
		flags.New(),
		fuzzyclock.New(),
		gapilotbuddy.New(),
		goodservice.New(),
		happyhour.New(),
		hurricanetracker.New(),
		hvvdepartures.New(),
		ifparank.New(),
		isitchristmas.New(),
		isstracker.New(),
		jokesjokeapi.New(),
		launchcountdown.New(),
		lirr.New(),
		mbta.New(),
		mbtanewtrains.New(),
		mlbleaders.New(),
		mlbscores.New(),
		mnlightrail.New(),
		moretransit.New(),
		moviequotes.New(),
		natdex.New(),
		nationaltoday.New(),
		nbascores.New(),
		nearearthobjs.New(),
		netatmo.New(),
		nflscores.New(),
		nft.New(),
		nhlnextgame.New(),
		nhlscores.New(),
		nightscout.New(),
		nixelclock.New(),
		noaabuoy.New(),
		nyancat.New(),
		nycbus.New(),
		ohhighwaysigns.New(),
		pagerduty.New(),
		pathtrainschedule.New(),
		phaseofmoon.New(),
		pokedex.New(),
		powerball.New(),
		purpleair.New(),
		randomslackmoji.New(),
		redditimages.New(),
		redditrplace.New(),
		sbbtimetable.New(),
		sfnextmuni.New(),
		snyk.New(),
		sportsscores.New(),
		sportsstandings.New(),
		spotthestation.New(),
		steam.New(),
		stepcounter.New(),
		strava.New(),
		sunrisesunset.New(),
		surflive.New(),
		tempest.New(),
		theysaidso.New(),
		todoist.New(),
		traffic.New(),
		transsee.New(),
		twitterfollows.New(),
		unsplash.New(),
		usgsearthquakes.New(),
		usyieldcurve.New(),
		vergetaglines.New(),
		verticalmessage.New(),
		wantedposter.New(),
		warframecycles.New(),
		weathermap.New(),
		whosthatpokemon.New(),
		worldclock.New(),
	}
}

// FindManifest finds an manifest at the given ID.
func FindManifest(id string) (*manifest.Manifest, error) {
	for _, app := range GetManifests() {
		if app.ID == id {
			return &app, nil
		}
	}

	return nil, errors.New("app manifest does not exist")
}