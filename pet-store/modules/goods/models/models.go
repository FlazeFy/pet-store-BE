package models

type (
	GetGoods struct {
		GoodsSlug     string `json:"goods_slug"`
		GoodsName     string `json:"goods_name"`
		GoodsDesc     string `json:"goods_desc"`
		GoodsCategory string `json:"goods_category"`
		GoodsImgUrl   string `json:"goods_img_url"`
		GoodsPrice    int    `json:"goods_price"`
		GoodsStock    int    `json:"goods_stock"`
	}
)
