package osm

import "fmt"

type OpenStreetMapQueryBuilder struct { }

func (osmQueryBuilder *OpenStreetMapQueryBuilder) GetDistrictIDFromLocation(lat float32, lon float32) string {

    return fmt.Sprintf(`SELECT gid FROM districts WHERE ST_DWithin(ST_SetSRID(ST_POINT(%f, %f),4326)::geography, the_geom,0);`, lat, lon)
}

