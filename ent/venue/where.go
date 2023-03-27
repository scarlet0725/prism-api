// Code generated by ent, DO NOT EDIT.

package venue

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/scarlet0725/prism-api/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Venue {
	return predicate.Venue(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Venue {
	return predicate.Venue(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Venue {
	return predicate.Venue(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Venue {
	return predicate.Venue(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Venue {
	return predicate.Venue(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Venue {
	return predicate.Venue(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Venue {
	return predicate.Venue(sql.FieldLTE(FieldID, id))
}

// VenueID applies equality check predicate on the "venue_id" field. It's identical to VenueIDEQ.
func VenueID(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldVenueID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldDescription, v))
}

// WebSite applies equality check predicate on the "web_site" field. It's identical to WebSiteEQ.
func WebSite(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldWebSite, v))
}

// Postcode applies equality check predicate on the "postcode" field. It's identical to PostcodeEQ.
func Postcode(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldPostcode, v))
}

// Prefecture applies equality check predicate on the "prefecture" field. It's identical to PrefectureEQ.
func Prefecture(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldPrefecture, v))
}

// City applies equality check predicate on the "city" field. It's identical to CityEQ.
func City(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldCity, v))
}

// Street applies equality check predicate on the "street" field. It's identical to StreetEQ.
func Street(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldStreet, v))
}

// IsOpen applies equality check predicate on the "is_open" field. It's identical to IsOpenEQ.
func IsOpen(v bool) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldIsOpen, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldDeletedAt, v))
}

// VenueIDEQ applies the EQ predicate on the "venue_id" field.
func VenueIDEQ(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldVenueID, v))
}

// VenueIDNEQ applies the NEQ predicate on the "venue_id" field.
func VenueIDNEQ(v string) predicate.Venue {
	return predicate.Venue(sql.FieldNEQ(FieldVenueID, v))
}

// VenueIDIn applies the In predicate on the "venue_id" field.
func VenueIDIn(vs ...string) predicate.Venue {
	return predicate.Venue(sql.FieldIn(FieldVenueID, vs...))
}

// VenueIDNotIn applies the NotIn predicate on the "venue_id" field.
func VenueIDNotIn(vs ...string) predicate.Venue {
	return predicate.Venue(sql.FieldNotIn(FieldVenueID, vs...))
}

// VenueIDGT applies the GT predicate on the "venue_id" field.
func VenueIDGT(v string) predicate.Venue {
	return predicate.Venue(sql.FieldGT(FieldVenueID, v))
}

// VenueIDGTE applies the GTE predicate on the "venue_id" field.
func VenueIDGTE(v string) predicate.Venue {
	return predicate.Venue(sql.FieldGTE(FieldVenueID, v))
}

// VenueIDLT applies the LT predicate on the "venue_id" field.
func VenueIDLT(v string) predicate.Venue {
	return predicate.Venue(sql.FieldLT(FieldVenueID, v))
}

// VenueIDLTE applies the LTE predicate on the "venue_id" field.
func VenueIDLTE(v string) predicate.Venue {
	return predicate.Venue(sql.FieldLTE(FieldVenueID, v))
}

// VenueIDContains applies the Contains predicate on the "venue_id" field.
func VenueIDContains(v string) predicate.Venue {
	return predicate.Venue(sql.FieldContains(FieldVenueID, v))
}

// VenueIDHasPrefix applies the HasPrefix predicate on the "venue_id" field.
func VenueIDHasPrefix(v string) predicate.Venue {
	return predicate.Venue(sql.FieldHasPrefix(FieldVenueID, v))
}

// VenueIDHasSuffix applies the HasSuffix predicate on the "venue_id" field.
func VenueIDHasSuffix(v string) predicate.Venue {
	return predicate.Venue(sql.FieldHasSuffix(FieldVenueID, v))
}

// VenueIDEqualFold applies the EqualFold predicate on the "venue_id" field.
func VenueIDEqualFold(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEqualFold(FieldVenueID, v))
}

// VenueIDContainsFold applies the ContainsFold predicate on the "venue_id" field.
func VenueIDContainsFold(v string) predicate.Venue {
	return predicate.Venue(sql.FieldContainsFold(FieldVenueID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Venue {
	return predicate.Venue(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Venue {
	return predicate.Venue(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Venue {
	return predicate.Venue(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Venue {
	return predicate.Venue(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Venue {
	return predicate.Venue(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Venue {
	return predicate.Venue(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Venue {
	return predicate.Venue(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Venue {
	return predicate.Venue(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Venue {
	return predicate.Venue(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Venue {
	return predicate.Venue(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Venue {
	return predicate.Venue(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Venue {
	return predicate.Venue(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Venue {
	return predicate.Venue(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Venue {
	return predicate.Venue(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Venue {
	return predicate.Venue(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Venue {
	return predicate.Venue(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Venue {
	return predicate.Venue(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Venue {
	return predicate.Venue(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Venue {
	return predicate.Venue(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Venue {
	return predicate.Venue(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Venue {
	return predicate.Venue(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Venue {
	return predicate.Venue(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Venue {
	return predicate.Venue(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Venue {
	return predicate.Venue(sql.FieldContainsFold(FieldDescription, v))
}

// WebSiteEQ applies the EQ predicate on the "web_site" field.
func WebSiteEQ(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldWebSite, v))
}

// WebSiteNEQ applies the NEQ predicate on the "web_site" field.
func WebSiteNEQ(v string) predicate.Venue {
	return predicate.Venue(sql.FieldNEQ(FieldWebSite, v))
}

// WebSiteIn applies the In predicate on the "web_site" field.
func WebSiteIn(vs ...string) predicate.Venue {
	return predicate.Venue(sql.FieldIn(FieldWebSite, vs...))
}

// WebSiteNotIn applies the NotIn predicate on the "web_site" field.
func WebSiteNotIn(vs ...string) predicate.Venue {
	return predicate.Venue(sql.FieldNotIn(FieldWebSite, vs...))
}

// WebSiteGT applies the GT predicate on the "web_site" field.
func WebSiteGT(v string) predicate.Venue {
	return predicate.Venue(sql.FieldGT(FieldWebSite, v))
}

// WebSiteGTE applies the GTE predicate on the "web_site" field.
func WebSiteGTE(v string) predicate.Venue {
	return predicate.Venue(sql.FieldGTE(FieldWebSite, v))
}

// WebSiteLT applies the LT predicate on the "web_site" field.
func WebSiteLT(v string) predicate.Venue {
	return predicate.Venue(sql.FieldLT(FieldWebSite, v))
}

// WebSiteLTE applies the LTE predicate on the "web_site" field.
func WebSiteLTE(v string) predicate.Venue {
	return predicate.Venue(sql.FieldLTE(FieldWebSite, v))
}

// WebSiteContains applies the Contains predicate on the "web_site" field.
func WebSiteContains(v string) predicate.Venue {
	return predicate.Venue(sql.FieldContains(FieldWebSite, v))
}

// WebSiteHasPrefix applies the HasPrefix predicate on the "web_site" field.
func WebSiteHasPrefix(v string) predicate.Venue {
	return predicate.Venue(sql.FieldHasPrefix(FieldWebSite, v))
}

// WebSiteHasSuffix applies the HasSuffix predicate on the "web_site" field.
func WebSiteHasSuffix(v string) predicate.Venue {
	return predicate.Venue(sql.FieldHasSuffix(FieldWebSite, v))
}

// WebSiteIsNil applies the IsNil predicate on the "web_site" field.
func WebSiteIsNil() predicate.Venue {
	return predicate.Venue(sql.FieldIsNull(FieldWebSite))
}

// WebSiteNotNil applies the NotNil predicate on the "web_site" field.
func WebSiteNotNil() predicate.Venue {
	return predicate.Venue(sql.FieldNotNull(FieldWebSite))
}

// WebSiteEqualFold applies the EqualFold predicate on the "web_site" field.
func WebSiteEqualFold(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEqualFold(FieldWebSite, v))
}

// WebSiteContainsFold applies the ContainsFold predicate on the "web_site" field.
func WebSiteContainsFold(v string) predicate.Venue {
	return predicate.Venue(sql.FieldContainsFold(FieldWebSite, v))
}

// PostcodeEQ applies the EQ predicate on the "postcode" field.
func PostcodeEQ(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldPostcode, v))
}

// PostcodeNEQ applies the NEQ predicate on the "postcode" field.
func PostcodeNEQ(v string) predicate.Venue {
	return predicate.Venue(sql.FieldNEQ(FieldPostcode, v))
}

// PostcodeIn applies the In predicate on the "postcode" field.
func PostcodeIn(vs ...string) predicate.Venue {
	return predicate.Venue(sql.FieldIn(FieldPostcode, vs...))
}

// PostcodeNotIn applies the NotIn predicate on the "postcode" field.
func PostcodeNotIn(vs ...string) predicate.Venue {
	return predicate.Venue(sql.FieldNotIn(FieldPostcode, vs...))
}

// PostcodeGT applies the GT predicate on the "postcode" field.
func PostcodeGT(v string) predicate.Venue {
	return predicate.Venue(sql.FieldGT(FieldPostcode, v))
}

// PostcodeGTE applies the GTE predicate on the "postcode" field.
func PostcodeGTE(v string) predicate.Venue {
	return predicate.Venue(sql.FieldGTE(FieldPostcode, v))
}

// PostcodeLT applies the LT predicate on the "postcode" field.
func PostcodeLT(v string) predicate.Venue {
	return predicate.Venue(sql.FieldLT(FieldPostcode, v))
}

// PostcodeLTE applies the LTE predicate on the "postcode" field.
func PostcodeLTE(v string) predicate.Venue {
	return predicate.Venue(sql.FieldLTE(FieldPostcode, v))
}

// PostcodeContains applies the Contains predicate on the "postcode" field.
func PostcodeContains(v string) predicate.Venue {
	return predicate.Venue(sql.FieldContains(FieldPostcode, v))
}

// PostcodeHasPrefix applies the HasPrefix predicate on the "postcode" field.
func PostcodeHasPrefix(v string) predicate.Venue {
	return predicate.Venue(sql.FieldHasPrefix(FieldPostcode, v))
}

// PostcodeHasSuffix applies the HasSuffix predicate on the "postcode" field.
func PostcodeHasSuffix(v string) predicate.Venue {
	return predicate.Venue(sql.FieldHasSuffix(FieldPostcode, v))
}

// PostcodeIsNil applies the IsNil predicate on the "postcode" field.
func PostcodeIsNil() predicate.Venue {
	return predicate.Venue(sql.FieldIsNull(FieldPostcode))
}

// PostcodeNotNil applies the NotNil predicate on the "postcode" field.
func PostcodeNotNil() predicate.Venue {
	return predicate.Venue(sql.FieldNotNull(FieldPostcode))
}

// PostcodeEqualFold applies the EqualFold predicate on the "postcode" field.
func PostcodeEqualFold(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEqualFold(FieldPostcode, v))
}

// PostcodeContainsFold applies the ContainsFold predicate on the "postcode" field.
func PostcodeContainsFold(v string) predicate.Venue {
	return predicate.Venue(sql.FieldContainsFold(FieldPostcode, v))
}

// PrefectureEQ applies the EQ predicate on the "prefecture" field.
func PrefectureEQ(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldPrefecture, v))
}

// PrefectureNEQ applies the NEQ predicate on the "prefecture" field.
func PrefectureNEQ(v string) predicate.Venue {
	return predicate.Venue(sql.FieldNEQ(FieldPrefecture, v))
}

// PrefectureIn applies the In predicate on the "prefecture" field.
func PrefectureIn(vs ...string) predicate.Venue {
	return predicate.Venue(sql.FieldIn(FieldPrefecture, vs...))
}

// PrefectureNotIn applies the NotIn predicate on the "prefecture" field.
func PrefectureNotIn(vs ...string) predicate.Venue {
	return predicate.Venue(sql.FieldNotIn(FieldPrefecture, vs...))
}

// PrefectureGT applies the GT predicate on the "prefecture" field.
func PrefectureGT(v string) predicate.Venue {
	return predicate.Venue(sql.FieldGT(FieldPrefecture, v))
}

// PrefectureGTE applies the GTE predicate on the "prefecture" field.
func PrefectureGTE(v string) predicate.Venue {
	return predicate.Venue(sql.FieldGTE(FieldPrefecture, v))
}

// PrefectureLT applies the LT predicate on the "prefecture" field.
func PrefectureLT(v string) predicate.Venue {
	return predicate.Venue(sql.FieldLT(FieldPrefecture, v))
}

// PrefectureLTE applies the LTE predicate on the "prefecture" field.
func PrefectureLTE(v string) predicate.Venue {
	return predicate.Venue(sql.FieldLTE(FieldPrefecture, v))
}

// PrefectureContains applies the Contains predicate on the "prefecture" field.
func PrefectureContains(v string) predicate.Venue {
	return predicate.Venue(sql.FieldContains(FieldPrefecture, v))
}

// PrefectureHasPrefix applies the HasPrefix predicate on the "prefecture" field.
func PrefectureHasPrefix(v string) predicate.Venue {
	return predicate.Venue(sql.FieldHasPrefix(FieldPrefecture, v))
}

// PrefectureHasSuffix applies the HasSuffix predicate on the "prefecture" field.
func PrefectureHasSuffix(v string) predicate.Venue {
	return predicate.Venue(sql.FieldHasSuffix(FieldPrefecture, v))
}

// PrefectureIsNil applies the IsNil predicate on the "prefecture" field.
func PrefectureIsNil() predicate.Venue {
	return predicate.Venue(sql.FieldIsNull(FieldPrefecture))
}

// PrefectureNotNil applies the NotNil predicate on the "prefecture" field.
func PrefectureNotNil() predicate.Venue {
	return predicate.Venue(sql.FieldNotNull(FieldPrefecture))
}

// PrefectureEqualFold applies the EqualFold predicate on the "prefecture" field.
func PrefectureEqualFold(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEqualFold(FieldPrefecture, v))
}

// PrefectureContainsFold applies the ContainsFold predicate on the "prefecture" field.
func PrefectureContainsFold(v string) predicate.Venue {
	return predicate.Venue(sql.FieldContainsFold(FieldPrefecture, v))
}

// CityEQ applies the EQ predicate on the "city" field.
func CityEQ(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldCity, v))
}

// CityNEQ applies the NEQ predicate on the "city" field.
func CityNEQ(v string) predicate.Venue {
	return predicate.Venue(sql.FieldNEQ(FieldCity, v))
}

// CityIn applies the In predicate on the "city" field.
func CityIn(vs ...string) predicate.Venue {
	return predicate.Venue(sql.FieldIn(FieldCity, vs...))
}

// CityNotIn applies the NotIn predicate on the "city" field.
func CityNotIn(vs ...string) predicate.Venue {
	return predicate.Venue(sql.FieldNotIn(FieldCity, vs...))
}

// CityGT applies the GT predicate on the "city" field.
func CityGT(v string) predicate.Venue {
	return predicate.Venue(sql.FieldGT(FieldCity, v))
}

// CityGTE applies the GTE predicate on the "city" field.
func CityGTE(v string) predicate.Venue {
	return predicate.Venue(sql.FieldGTE(FieldCity, v))
}

// CityLT applies the LT predicate on the "city" field.
func CityLT(v string) predicate.Venue {
	return predicate.Venue(sql.FieldLT(FieldCity, v))
}

// CityLTE applies the LTE predicate on the "city" field.
func CityLTE(v string) predicate.Venue {
	return predicate.Venue(sql.FieldLTE(FieldCity, v))
}

// CityContains applies the Contains predicate on the "city" field.
func CityContains(v string) predicate.Venue {
	return predicate.Venue(sql.FieldContains(FieldCity, v))
}

// CityHasPrefix applies the HasPrefix predicate on the "city" field.
func CityHasPrefix(v string) predicate.Venue {
	return predicate.Venue(sql.FieldHasPrefix(FieldCity, v))
}

// CityHasSuffix applies the HasSuffix predicate on the "city" field.
func CityHasSuffix(v string) predicate.Venue {
	return predicate.Venue(sql.FieldHasSuffix(FieldCity, v))
}

// CityIsNil applies the IsNil predicate on the "city" field.
func CityIsNil() predicate.Venue {
	return predicate.Venue(sql.FieldIsNull(FieldCity))
}

// CityNotNil applies the NotNil predicate on the "city" field.
func CityNotNil() predicate.Venue {
	return predicate.Venue(sql.FieldNotNull(FieldCity))
}

// CityEqualFold applies the EqualFold predicate on the "city" field.
func CityEqualFold(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEqualFold(FieldCity, v))
}

// CityContainsFold applies the ContainsFold predicate on the "city" field.
func CityContainsFold(v string) predicate.Venue {
	return predicate.Venue(sql.FieldContainsFold(FieldCity, v))
}

// StreetEQ applies the EQ predicate on the "street" field.
func StreetEQ(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldStreet, v))
}

// StreetNEQ applies the NEQ predicate on the "street" field.
func StreetNEQ(v string) predicate.Venue {
	return predicate.Venue(sql.FieldNEQ(FieldStreet, v))
}

// StreetIn applies the In predicate on the "street" field.
func StreetIn(vs ...string) predicate.Venue {
	return predicate.Venue(sql.FieldIn(FieldStreet, vs...))
}

// StreetNotIn applies the NotIn predicate on the "street" field.
func StreetNotIn(vs ...string) predicate.Venue {
	return predicate.Venue(sql.FieldNotIn(FieldStreet, vs...))
}

// StreetGT applies the GT predicate on the "street" field.
func StreetGT(v string) predicate.Venue {
	return predicate.Venue(sql.FieldGT(FieldStreet, v))
}

// StreetGTE applies the GTE predicate on the "street" field.
func StreetGTE(v string) predicate.Venue {
	return predicate.Venue(sql.FieldGTE(FieldStreet, v))
}

// StreetLT applies the LT predicate on the "street" field.
func StreetLT(v string) predicate.Venue {
	return predicate.Venue(sql.FieldLT(FieldStreet, v))
}

// StreetLTE applies the LTE predicate on the "street" field.
func StreetLTE(v string) predicate.Venue {
	return predicate.Venue(sql.FieldLTE(FieldStreet, v))
}

// StreetContains applies the Contains predicate on the "street" field.
func StreetContains(v string) predicate.Venue {
	return predicate.Venue(sql.FieldContains(FieldStreet, v))
}

// StreetHasPrefix applies the HasPrefix predicate on the "street" field.
func StreetHasPrefix(v string) predicate.Venue {
	return predicate.Venue(sql.FieldHasPrefix(FieldStreet, v))
}

// StreetHasSuffix applies the HasSuffix predicate on the "street" field.
func StreetHasSuffix(v string) predicate.Venue {
	return predicate.Venue(sql.FieldHasSuffix(FieldStreet, v))
}

// StreetIsNil applies the IsNil predicate on the "street" field.
func StreetIsNil() predicate.Venue {
	return predicate.Venue(sql.FieldIsNull(FieldStreet))
}

// StreetNotNil applies the NotNil predicate on the "street" field.
func StreetNotNil() predicate.Venue {
	return predicate.Venue(sql.FieldNotNull(FieldStreet))
}

// StreetEqualFold applies the EqualFold predicate on the "street" field.
func StreetEqualFold(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEqualFold(FieldStreet, v))
}

// StreetContainsFold applies the ContainsFold predicate on the "street" field.
func StreetContainsFold(v string) predicate.Venue {
	return predicate.Venue(sql.FieldContainsFold(FieldStreet, v))
}

// IsOpenEQ applies the EQ predicate on the "is_open" field.
func IsOpenEQ(v bool) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldIsOpen, v))
}

// IsOpenNEQ applies the NEQ predicate on the "is_open" field.
func IsOpenNEQ(v bool) predicate.Venue {
	return predicate.Venue(sql.FieldNEQ(FieldIsOpen, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Venue {
	return predicate.Venue(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Venue {
	return predicate.Venue(sql.FieldNotNull(FieldDeletedAt))
}

// HasEvents applies the HasEdge predicate on the "events" edge.
func HasEvents() predicate.Venue {
	return predicate.Venue(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, EventsTable, EventsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEventsWith applies the HasEdge predicate on the "events" edge with a given conditions (other predicates).
func HasEventsWith(preds ...predicate.Event) predicate.Venue {
	return predicate.Venue(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EventsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, EventsTable, EventsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Venue) predicate.Venue {
	return predicate.Venue(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Venue) predicate.Venue {
	return predicate.Venue(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Venue) predicate.Venue {
	return predicate.Venue(func(s *sql.Selector) {
		p(s.Not())
	})
}
