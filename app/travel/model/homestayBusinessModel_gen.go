// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"time"

	"study-zero/pkg/globalkey"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	homestayBusinessFieldNames          = builder.RawFieldNames(&HomestayBusiness{})
	homestayBusinessRows                = strings.Join(homestayBusinessFieldNames, ",")
	homestayBusinessRowsExpectAutoSet   = strings.Join(stringx.Remove(homestayBusinessFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	homestayBusinessRowsWithPlaceHolder = strings.Join(stringx.Remove(homestayBusinessFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheLooklookTravelHomestayBusinessIdPrefix     = "cache:looklookTravel:homestayBusiness:id:"
	cacheLooklookTravelHomestayBusinessUserIdPrefix = "cache:looklookTravel:homestayBusiness:userId:"
)

type (
	homestayBusinessModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *HomestayBusiness) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*HomestayBusiness, error)
		FindOneByUserId(ctx context.Context, userId int64) (*HomestayBusiness, error)
		Update(ctx context.Context, session sqlx.Session, data *HomestayBusiness) (sql.Result, error)
		UpdateWithVersion(ctx context.Context, session sqlx.Session, data *HomestayBusiness) error
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		SelectBuilder() squirrel.SelectBuilder
		DeleteSoft(ctx context.Context, session sqlx.Session, data *HomestayBusiness) error
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder, field string) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder, field string) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*HomestayBusiness, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*HomestayBusiness, error)
		FindPageListByPageWithTotal(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*HomestayBusiness, int64, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*HomestayBusiness, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*HomestayBusiness, error)
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultHomestayBusinessModel struct {
		sqlc.CachedConn
		table string
	}

	HomestayBusiness struct {
		Id          int64     `db:"id"`
		CreateTime  time.Time `db:"create_time"`
		UpdateTime  time.Time `db:"update_time"`
		DeleteTime  time.Time `db:"delete_time"`
		DelState    int64     `db:"del_state"`
		Title       string    `db:"title"`        // 店铺名称
		UserId      int64     `db:"user_id"`      // 关联的用户id
		Info        string    `db:"info"`         // 店铺介绍
		BossInfo    string    `db:"boss_info"`    // 房东介绍
		LicenseFron string    `db:"license_fron"` // 营业执照正面
		LicenseBack string    `db:"license_back"` // 营业执照背面
		RowState    int64     `db:"row_state"`    // 0:禁止营业 1:正常营业
		Star        float64   `db:"star"`         // 店铺整体评价，冗余
		Tags        string    `db:"tags"`         // 每个店家一个标签，自己编辑
		Cover       string    `db:"cover"`        // 封面图
		HeaderImg   string    `db:"header_img"`   // 店招门头图片
		Version     int64     `db:"version"`      // 版本号
	}
)

func newHomestayBusinessModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultHomestayBusinessModel {
	return &defaultHomestayBusinessModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`homestay_business`",
	}
}

func (m *defaultHomestayBusinessModel) Insert(ctx context.Context, session sqlx.Session, data *HomestayBusiness) (sql.Result, error) {
	data.DeleteTime = time.Unix(0, 0)
	data.DelState = globalkey.DelStateNo
	looklookTravelHomestayBusinessIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelHomestayBusinessIdPrefix, data.Id)
	looklookTravelHomestayBusinessUserIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelHomestayBusinessUserIdPrefix, data.UserId)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, homestayBusinessRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Title, data.UserId, data.Info, data.BossInfo, data.LicenseFron, data.LicenseBack, data.RowState, data.Star, data.Tags, data.Cover, data.HeaderImg, data.Version)
		}
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Title, data.UserId, data.Info, data.BossInfo, data.LicenseFron, data.LicenseBack, data.RowState, data.Star, data.Tags, data.Cover, data.HeaderImg, data.Version)
	}, looklookTravelHomestayBusinessIdKey, looklookTravelHomestayBusinessUserIdKey)
}

func (m *defaultHomestayBusinessModel) FindOne(ctx context.Context, id int64) (*HomestayBusiness, error) {
	looklookTravelHomestayBusinessIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelHomestayBusinessIdPrefix, id)
	var resp HomestayBusiness
	err := m.QueryRowCtx(ctx, &resp, looklookTravelHomestayBusinessIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", homestayBusinessRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id, globalkey.DelStateNo)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultHomestayBusinessModel) FindOneByUserId(ctx context.Context, userId int64) (*HomestayBusiness, error) {
	looklookTravelHomestayBusinessUserIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelHomestayBusinessUserIdPrefix, userId)
	var resp HomestayBusiness
	err := m.QueryRowIndexCtx(ctx, &resp, looklookTravelHomestayBusinessUserIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? and del_state = ? limit 1", homestayBusinessRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userId, globalkey.DelStateNo); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultHomestayBusinessModel) Update(ctx context.Context, session sqlx.Session, newData *HomestayBusiness) (sql.Result, error) {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return nil, err
	}
	looklookTravelHomestayBusinessIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelHomestayBusinessIdPrefix, data.Id)
	looklookTravelHomestayBusinessUserIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelHomestayBusinessUserIdPrefix, data.UserId)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, homestayBusinessRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, newData.DeleteTime, newData.DelState, newData.Title, newData.UserId, newData.Info, newData.BossInfo, newData.LicenseFron, newData.LicenseBack, newData.RowState, newData.Star, newData.Tags, newData.Cover, newData.HeaderImg, newData.Version, newData.Id)
		}
		return conn.ExecCtx(ctx, query, newData.DeleteTime, newData.DelState, newData.Title, newData.UserId, newData.Info, newData.BossInfo, newData.LicenseFron, newData.LicenseBack, newData.RowState, newData.Star, newData.Tags, newData.Cover, newData.HeaderImg, newData.Version, newData.Id)
	}, looklookTravelHomestayBusinessIdKey, looklookTravelHomestayBusinessUserIdKey)
}

func (m *defaultHomestayBusinessModel) UpdateWithVersion(ctx context.Context, session sqlx.Session, newData *HomestayBusiness) error {

	oldVersion := newData.Version
	newData.Version += 1

	var sqlResult sql.Result
	var err error

	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}
	looklookTravelHomestayBusinessIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelHomestayBusinessIdPrefix, data.Id)
	looklookTravelHomestayBusinessUserIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelHomestayBusinessUserIdPrefix, data.UserId)
	sqlResult, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and version = ? ", m.table, homestayBusinessRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, newData.DeleteTime, newData.DelState, newData.Title, newData.UserId, newData.Info, newData.BossInfo, newData.LicenseFron, newData.LicenseBack, newData.RowState, newData.Star, newData.Tags, newData.Cover, newData.HeaderImg, newData.Version, newData.Id, oldVersion)
		}
		return conn.ExecCtx(ctx, query, newData.DeleteTime, newData.DelState, newData.Title, newData.UserId, newData.Info, newData.BossInfo, newData.LicenseFron, newData.LicenseBack, newData.RowState, newData.Star, newData.Tags, newData.Cover, newData.HeaderImg, newData.Version, newData.Id, oldVersion)
	}, looklookTravelHomestayBusinessIdKey, looklookTravelHomestayBusinessUserIdKey)
	if err != nil {
		return err
	}
	updateCount, err := sqlResult.RowsAffected()
	if err != nil {
		return err
	}
	if updateCount == 0 {
		return ErrNoRowsUpdate
	}

	return nil
}

func (m *defaultHomestayBusinessModel) DeleteSoft(ctx context.Context, session sqlx.Session, data *HomestayBusiness) error {
	data.DelState = globalkey.DelStateYes
	data.DeleteTime = time.Now()
	if err := m.UpdateWithVersion(ctx, session, data); err != nil {
		return errors.Wrapf(errors.New("delete soft failed "), "HomestayBusinessModel delete err : %+v", err)
	}
	return nil
}

func (m *defaultHomestayBusinessModel) FindSum(ctx context.Context, builder squirrel.SelectBuilder, field string) (float64, error) {

	if len(field) == 0 {
		return 0, errors.Wrapf(errors.New("FindSum Least One Field"), "FindSum Least One Field")
	}

	builder = builder.Columns("IFNULL(SUM(" + field + "),0)")

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return 0, err
	}

	var resp float64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultHomestayBusinessModel) FindCount(ctx context.Context, builder squirrel.SelectBuilder, field string) (int64, error) {

	if len(field) == 0 {
		return 0, errors.Wrapf(errors.New("FindCount Least One Field"), "FindCount Least One Field")
	}

	builder = builder.Columns("COUNT(" + field + ")")

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultHomestayBusinessModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*HomestayBusiness, error) {

	builder = builder.Columns(homestayBusinessRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*HomestayBusiness
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultHomestayBusinessModel) FindPageListByPage(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*HomestayBusiness, error) {

	builder = builder.Columns(homestayBusinessRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*HomestayBusiness
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultHomestayBusinessModel) FindPageListByPageWithTotal(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*HomestayBusiness, int64, error) {

	total, err := m.FindCount(ctx, builder, "id")
	if err != nil {
		return nil, 0, err
	}

	builder = builder.Columns(homestayBusinessRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, total, err
	}

	var resp []*HomestayBusiness
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, total, nil
	default:
		return nil, total, err
	}
}

func (m *defaultHomestayBusinessModel) FindPageListByIdDESC(ctx context.Context, builder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*HomestayBusiness, error) {

	builder = builder.Columns(homestayBusinessRows)

	if preMinId > 0 {
		builder = builder.Where(" id < ? ", preMinId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*HomestayBusiness
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultHomestayBusinessModel) FindPageListByIdASC(ctx context.Context, builder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*HomestayBusiness, error) {

	builder = builder.Columns(homestayBusinessRows)

	if preMaxId > 0 {
		builder = builder.Where(" id > ? ", preMaxId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*HomestayBusiness
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultHomestayBusinessModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

func (m *defaultHomestayBusinessModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}
func (m *defaultHomestayBusinessModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	looklookTravelHomestayBusinessIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelHomestayBusinessIdPrefix, id)
	looklookTravelHomestayBusinessUserIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelHomestayBusinessUserIdPrefix, data.UserId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, looklookTravelHomestayBusinessIdKey, looklookTravelHomestayBusinessUserIdKey)
	return err
}
func (m *defaultHomestayBusinessModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheLooklookTravelHomestayBusinessIdPrefix, primary)
}
func (m *defaultHomestayBusinessModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", homestayBusinessRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary, globalkey.DelStateNo)
}

func (m *defaultHomestayBusinessModel) tableName() string {
	return m.table
}
