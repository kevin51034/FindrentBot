package main

type FlexMessage struct {
	Type   string `json:"type"`
	Header struct {
		Type     string `json:"type"`
		Layout   string `json:"layout"`
		Contents []struct {
			Type     string `json:"type"`
			Layout   string `json:"layout"`
			Contents []struct {
				Type        string `json:"type"`
				URL         string `json:"url,omitempty"`
				Size        string `json:"size,omitempty"`
				AspectMode  string `json:"aspectMode,omitempty"`
				AspectRatio string `json:"aspectRatio,omitempty"`
				Gravity     string `json:"gravity,omitempty"`
				Flex        int    `json:"flex"`
				Layout      string `json:"layout,omitempty"`
				Contents    []struct {
					Type    string `json:"type"`
					Text    string `json:"text"`
					Size    string `json:"size"`
					Color   string `json:"color"`
					Align   string `json:"align"`
					Gravity string `json:"gravity"`
				} `json:"contents,omitempty"`
				BackgroundColor string `json:"backgroundColor,omitempty"`
				PaddingAll      string `json:"paddingAll,omitempty"`
				PaddingStart    string `json:"paddingStart,omitempty"`
				PaddingEnd      string `json:"paddingEnd,omitempty"`
				Position        string `json:"position,omitempty"`
				OffsetStart     string `json:"offsetStart,omitempty"`
				OffsetTop       string `json:"offsetTop,omitempty"`
				CornerRadius    string `json:"cornerRadius,omitempty"`
				Width           string `json:"width,omitempty"`
				Height          string `json:"height,omitempty"`
			} `json:"contents"`
		} `json:"contents"`
		PaddingAll string `json:"paddingAll"`
	} `json:"header"`
	Body struct {
		Type     string `json:"type"`
		Layout   string `json:"layout"`
		Contents []struct {
			Type     string `json:"type"`
			Layout   string `json:"layout"`
			Contents []struct {
				Type     string `json:"type"`
				Layout   string `json:"layout"`
				Contents []struct {
					Type     string        `json:"type"`
					Contents []interface{} `json:"contents,omitempty"`
					Size     string        `json:"size"`
					Wrap     bool          `json:"wrap,omitempty"`
					Text     string        `json:"text"`
					Color    string        `json:"color"`
					Weight   string        `json:"weight,omitempty"`
				} `json:"contents"`
				Spacing         string `json:"spacing,omitempty"`
				PaddingAll      string `json:"paddingAll,omitempty"`
				BackgroundColor string `json:"backgroundColor,omitempty"`
				CornerRadius    string `json:"cornerRadius,omitempty"`
				Margin          string `json:"margin,omitempty"`
			} `json:"contents"`
		} `json:"contents"`
		PaddingAll      string `json:"paddingAll"`
		BackgroundColor string `json:"backgroundColor"`
	} `json:"body"`
	Action struct {
		Type  string `json:"type"`
		Label string `json:"label"`
		URI   string `json:"uri"`
	} `json:"action"`
}

func NewJSONData() []byte {
	jsonData := []byte(`{
		"type": "bubble",
		"header": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "box",
			  "layout": "horizontal",
			  "contents": [
				{
				  "type": "image",
				  "url": "https://scdn.line-apps.com/n/channel_devcenter/img/flexsnapshot/clip/clip4.jpg",
				  "size": "full",
				  "aspectMode": "cover",
				  "aspectRatio": "150:196",
				  "gravity": "center",
				  "flex": 1
				},
				{
				  "type": "box",
				  "layout": "horizontal",
				  "contents": [
					{
					  "type": "text",
					  "text": "NEW",
					  "size": "xs",
					  "color": "#ffffff",
					  "align": "center",
					  "gravity": "center"
					}
				  ],
				  "backgroundColor": "#EC3D44",
				  "paddingAll": "2px",
				  "paddingStart": "4px",
				  "paddingEnd": "4px",
				  "flex": 0,
				  "position": "absolute",
				  "offsetStart": "18px",
				  "offsetTop": "18px",
				  "cornerRadius": "100px",
				  "width": "48px",
				  "height": "25px"
				}
			  ]
			}
		  ],
		  "paddingAll": "0px"
		},
		"body": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "box",
			  "layout": "vertical",
			  "contents": [
				{
				  "type": "box",
				  "layout": "vertical",
				  "contents": [
					{
					  "type": "text",
					  "contents": [],
					  "size": "xl",
					  "wrap": true,
					  "text": "Cony Residence",
					  "color": "#ffffff",
					  "weight": "bold"
					},
					{
					  "type": "text",
					  "text": "3 Bedrooms, ¥35,000",
					  "color": "#ffffffcc",
					  "size": "sm"
					}
				  ],
				  "spacing": "sm"
				},
				{
				  "type": "box",
				  "layout": "vertical",
				  "contents": [
					{
						  "type": "text",
						  "size": "sm",
						  "wrap": true,
						  "margin": "lg",
						  "color": "#ffffffde",
						  "text": "Private Pool, Delivery box, Floor heating, Private Cinema"
					}
				  ],
				  "paddingAll": "13px",
				  "backgroundColor": "#ffffff1A",
				  "cornerRadius": "2px",
				  "margin": "xl"
				}
			  ]
			}
		  ],
		  "paddingAll": "20px",
		  "backgroundColor": "#464F69"
		},
		"action": {
			"type": "uri",
			"label": "action",
			"uri": "http://linecorp.com/"
		}
	  }`)
	
	return jsonData
}
