// Code generated by ent, DO NOT EDIT.

package tramite

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/francoganga/go_reno2/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// AnoPresupuestario applies equality check predicate on the "anoPresupuestario" field. It's identical to AnoPresupuestarioEQ.
func AnoPresupuestario(v int) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAnoPresupuestario), v))
	})
}

// Link applies equality check predicate on the "link" field. It's identical to LinkEQ.
func Link(v string) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLink), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Categoria applies equality check predicate on the "categoria" field. It's identical to CategoriaEQ.
func Categoria(v string) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCategoria), v))
	})
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v int) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVersion), v))
	})
}

// AnoPresupuestarioEQ applies the EQ predicate on the "anoPresupuestario" field.
func AnoPresupuestarioEQ(v int) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAnoPresupuestario), v))
	})
}

// AnoPresupuestarioNEQ applies the NEQ predicate on the "anoPresupuestario" field.
func AnoPresupuestarioNEQ(v int) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAnoPresupuestario), v))
	})
}

// AnoPresupuestarioIn applies the In predicate on the "anoPresupuestario" field.
func AnoPresupuestarioIn(vs ...int) predicate.Tramite {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAnoPresupuestario), v...))
	})
}

// AnoPresupuestarioNotIn applies the NotIn predicate on the "anoPresupuestario" field.
func AnoPresupuestarioNotIn(vs ...int) predicate.Tramite {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAnoPresupuestario), v...))
	})
}

// AnoPresupuestarioGT applies the GT predicate on the "anoPresupuestario" field.
func AnoPresupuestarioGT(v int) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAnoPresupuestario), v))
	})
}

// AnoPresupuestarioGTE applies the GTE predicate on the "anoPresupuestario" field.
func AnoPresupuestarioGTE(v int) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAnoPresupuestario), v))
	})
}

// AnoPresupuestarioLT applies the LT predicate on the "anoPresupuestario" field.
func AnoPresupuestarioLT(v int) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAnoPresupuestario), v))
	})
}

// AnoPresupuestarioLTE applies the LTE predicate on the "anoPresupuestario" field.
func AnoPresupuestarioLTE(v int) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAnoPresupuestario), v))
	})
}

// LinkEQ applies the EQ predicate on the "link" field.
func LinkEQ(v string) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLink), v))
	})
}

// LinkNEQ applies the NEQ predicate on the "link" field.
func LinkNEQ(v string) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLink), v))
	})
}

// LinkIn applies the In predicate on the "link" field.
func LinkIn(vs ...string) predicate.Tramite {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLink), v...))
	})
}

// LinkNotIn applies the NotIn predicate on the "link" field.
func LinkNotIn(vs ...string) predicate.Tramite {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLink), v...))
	})
}

// LinkGT applies the GT predicate on the "link" field.
func LinkGT(v string) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLink), v))
	})
}

// LinkGTE applies the GTE predicate on the "link" field.
func LinkGTE(v string) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLink), v))
	})
}

// LinkLT applies the LT predicate on the "link" field.
func LinkLT(v string) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLink), v))
	})
}

// LinkLTE applies the LTE predicate on the "link" field.
func LinkLTE(v string) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLink), v))
	})
}

// LinkContains applies the Contains predicate on the "link" field.
func LinkContains(v string) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLink), v))
	})
}

// LinkHasPrefix applies the HasPrefix predicate on the "link" field.
func LinkHasPrefix(v string) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLink), v))
	})
}

// LinkHasSuffix applies the HasSuffix predicate on the "link" field.
func LinkHasSuffix(v string) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLink), v))
	})
}

// LinkEqualFold applies the EqualFold predicate on the "link" field.
func LinkEqualFold(v string) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLink), v))
	})
}

// LinkContainsFold applies the ContainsFold predicate on the "link" field.
func LinkContainsFold(v string) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLink), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Tramite {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Tramite {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Tramite {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Tramite {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// CategoriaEQ applies the EQ predicate on the "categoria" field.
func CategoriaEQ(v string) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCategoria), v))
	})
}

// CategoriaNEQ applies the NEQ predicate on the "categoria" field.
func CategoriaNEQ(v string) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCategoria), v))
	})
}

// CategoriaIn applies the In predicate on the "categoria" field.
func CategoriaIn(vs ...string) predicate.Tramite {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCategoria), v...))
	})
}

// CategoriaNotIn applies the NotIn predicate on the "categoria" field.
func CategoriaNotIn(vs ...string) predicate.Tramite {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCategoria), v...))
	})
}

// CategoriaGT applies the GT predicate on the "categoria" field.
func CategoriaGT(v string) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCategoria), v))
	})
}

// CategoriaGTE applies the GTE predicate on the "categoria" field.
func CategoriaGTE(v string) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCategoria), v))
	})
}

// CategoriaLT applies the LT predicate on the "categoria" field.
func CategoriaLT(v string) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCategoria), v))
	})
}

// CategoriaLTE applies the LTE predicate on the "categoria" field.
func CategoriaLTE(v string) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCategoria), v))
	})
}

// CategoriaContains applies the Contains predicate on the "categoria" field.
func CategoriaContains(v string) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCategoria), v))
	})
}

// CategoriaHasPrefix applies the HasPrefix predicate on the "categoria" field.
func CategoriaHasPrefix(v string) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCategoria), v))
	})
}

// CategoriaHasSuffix applies the HasSuffix predicate on the "categoria" field.
func CategoriaHasSuffix(v string) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCategoria), v))
	})
}

// CategoriaEqualFold applies the EqualFold predicate on the "categoria" field.
func CategoriaEqualFold(v string) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCategoria), v))
	})
}

// CategoriaContainsFold applies the ContainsFold predicate on the "categoria" field.
func CategoriaContainsFold(v string) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCategoria), v))
	})
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v int) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVersion), v))
	})
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v int) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldVersion), v))
	})
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...int) predicate.Tramite {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldVersion), v...))
	})
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...int) predicate.Tramite {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldVersion), v...))
	})
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v int) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldVersion), v))
	})
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v int) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldVersion), v))
	})
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v int) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldVersion), v))
	})
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v int) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldVersion), v))
	})
}

// HasEvents applies the HasEdge predicate on the "events" edge.
func HasEvents() predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EventsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EventsTable, EventsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEventsWith applies the HasEdge predicate on the "events" edge with a given conditions (other predicates).
func HasEventsWith(preds ...predicate.Event) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EventsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EventsTable, EventsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Tramite) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Tramite) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
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
func Not(p predicate.Tramite) predicate.Tramite {
	return predicate.Tramite(func(s *sql.Selector) {
		p(s.Not())
	})
}
