// MIT License

// Copyright (c) 2022 Tree Xie

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package charts

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

func TestTableChart(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		render func(*Painter) ([]byte, error)
		result string
	}{
		{
			render: func(p *Painter) ([]byte, error) {
				_, err := NewTableChart(p, TableChartOption{
					Header: []string{
						"Name",
						"Age",
						"Address",
						"Tag",
						"Action",
					},
					Spans: []int{
						1,
						1,
						2,
						1,
						// span和header不匹配，最后自动设置为1
						// 1,
					},
					Data: [][]string{
						{
							"John Brown",
							"32",
							"New York No. 1 Lake Park",
							"nice, developer",
							"Send Mail",
						},
						{
							"Jim Green	",
							"42",
							"London No. 1 Lake Park",
							"wow",
							"Send Mail",
						},
						{
							"Joe Black	",
							"32",
							"Sidney No. 1 Lake Park",
							"cool, teacher",
							"Send Mail",
						},
					},
				}).Render()
				if err != nil {
					return nil, err
				}
				return p.Bytes()
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"600\" height=\"400\">\\n<path  d=\"M 0 0\nL 600 0\nL 600 35\nL 0 35\nL 0 0\" style=\"stroke-width:0;stroke:none;fill:rgba(240,240,240,1.0)\"/><path  d=\"M 0 35\nL 600 35\nL 600 90\nL 0 90\nL 0 35\" style=\"stroke-width:0;stroke:none;fill:rgba(255,255,255,1.0)\"/><path  d=\"M 0 90\nL 600 90\nL 600 125\nL 0 125\nL 0 90\" style=\"stroke-width:0;stroke:none;fill:rgba(247,247,247,1.0)\"/><path  d=\"M 0 125\nL 600 125\nL 600 180\nL 0 180\nL 0 125\" style=\"stroke-width:0;stroke:none;fill:rgba(255,255,255,1.0)\"/><text x=\"10\" y=\"22\" style=\"stroke-width:0;stroke:none;fill:rgba(98,105,118,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Name</text><text x=\"110\" y=\"22\" style=\"stroke-width:0;stroke:none;fill:rgba(98,105,118,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Age</text><text x=\"210\" y=\"22\" style=\"stroke-width:0;stroke:none;fill:rgba(98,105,118,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Address</text><text x=\"410\" y=\"22\" style=\"stroke-width:0;stroke:none;fill:rgba(98,105,118,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Tag</text><text x=\"510\" y=\"22\" style=\"stroke-width:0;stroke:none;fill:rgba(98,105,118,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Action</text><text x=\"10\" y=\"57\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">John</text><text x=\"10\" y=\"77\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Brown</text><text x=\"110\" y=\"57\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">32</text><text x=\"210\" y=\"57\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">New York No. 1 Lake Park</text><text x=\"410\" y=\"57\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">nice,</text><text x=\"410\" y=\"77\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">developer</text><text x=\"510\" y=\"57\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Send Mail</text><text x=\"10\" y=\"112\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Jim Green</text><text x=\"110\" y=\"112\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">42</text><text x=\"210\" y=\"112\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">London No. 1 Lake Park</text><text x=\"410\" y=\"112\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">wow</text><text x=\"510\" y=\"112\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Send Mail</text><text x=\"10\" y=\"147\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Joe Black</text><text x=\"110\" y=\"147\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">32</text><text x=\"210\" y=\"147\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Sidney No. 1 Lake Park</text><text x=\"410\" y=\"147\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">cool,</text><text x=\"410\" y=\"167\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">teacher</text><text x=\"510\" y=\"147\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Send Mail</text></svg>",
		},
		{
			render: func(p *Painter) ([]byte, error) {
				_, err := NewTableChart(p, TableChartOption{
					Header: []string{
						"Name",
						"Age",
						"Address",
						"Tag",
						"Action",
					},
					Data: [][]string{
						{
							"John Brown",
							"32",
							"New York No. 1 Lake Park",
							"nice, developer",
							"Send Mail",
						},
						{
							"Jim Green	",
							"42",
							"London No. 1 Lake Park",
							"wow",
							"Send Mail",
						},
						{
							"Joe Black	",
							"32",
							"Sidney No. 1 Lake Park",
							"cool, teacher",
							"Send Mail",
						},
					},
				}).Render()
				if err != nil {
					return nil, err
				}
				return p.Bytes()
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"600\" height=\"400\">\\n<path  d=\"M 0 0\nL 600 0\nL 600 35\nL 0 35\nL 0 0\" style=\"stroke-width:0;stroke:none;fill:rgba(240,240,240,1.0)\"/><path  d=\"M 0 35\nL 600 35\nL 600 90\nL 0 90\nL 0 35\" style=\"stroke-width:0;stroke:none;fill:rgba(255,255,255,1.0)\"/><path  d=\"M 0 90\nL 600 90\nL 600 145\nL 0 145\nL 0 90\" style=\"stroke-width:0;stroke:none;fill:rgba(247,247,247,1.0)\"/><path  d=\"M 0 145\nL 600 145\nL 600 200\nL 0 200\nL 0 145\" style=\"stroke-width:0;stroke:none;fill:rgba(255,255,255,1.0)\"/><text x=\"10\" y=\"22\" style=\"stroke-width:0;stroke:none;fill:rgba(98,105,118,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Name</text><text x=\"130\" y=\"22\" style=\"stroke-width:0;stroke:none;fill:rgba(98,105,118,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Age</text><text x=\"250\" y=\"22\" style=\"stroke-width:0;stroke:none;fill:rgba(98,105,118,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Address</text><text x=\"370\" y=\"22\" style=\"stroke-width:0;stroke:none;fill:rgba(98,105,118,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Tag</text><text x=\"490\" y=\"22\" style=\"stroke-width:0;stroke:none;fill:rgba(98,105,118,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Action</text><text x=\"10\" y=\"57\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">John Brown</text><text x=\"130\" y=\"57\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">32</text><text x=\"250\" y=\"57\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">New York No.</text><text x=\"250\" y=\"77\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">1 Lake Park</text><text x=\"370\" y=\"57\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">nice,</text><text x=\"370\" y=\"77\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">developer</text><text x=\"490\" y=\"57\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Send Mail</text><text x=\"10\" y=\"112\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Jim Green</text><text x=\"130\" y=\"112\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">42</text><text x=\"250\" y=\"112\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">London No. 1</text><text x=\"250\" y=\"132\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Lake Park</text><text x=\"370\" y=\"112\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">wow</text><text x=\"490\" y=\"112\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Send Mail</text><text x=\"10\" y=\"167\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Joe Black</text><text x=\"130\" y=\"167\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">32</text><text x=\"250\" y=\"167\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Sidney No. 1</text><text x=\"250\" y=\"187\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Lake Park</text><text x=\"370\" y=\"167\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">cool, teacher</text><text x=\"490\" y=\"167\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Send Mail</text></svg>",
		},
		{
			render: func(p *Painter) ([]byte, error) {
				_, err := NewTableChart(p, TableChartOption{
					Header: []string{
						"Name",
						"Age",
						"Address",
						"Tag",
						"Action",
					},
					Data: [][]string{
						{
							"John Brown",
							"32",
							"New York No. 1 Lake Park",
							"nice, developer",
							"Send Mail",
						},
						{
							"Jim Green	",
							"42",
							"London No. 1 Lake Park",
							"wow",
							"Send Mail",
						},
						{
							"Joe Black	",
							"32",
							"Sidney No. 1 Lake Park",
							"cool, teacher",
							"Send Mail",
						},
					},
					RowSpans: map[int][]CellSpan{
						4: {
							{
								RowFrom: 0,
								RowTo:   1,
							},
						},
					},
					StrokeColor:          drawing.ColorBlack,
					StrokeWidth:          1,
					DisableRowBackground: true,
				}).Render()
				if err != nil {
					return nil, err
				}
				return p.Bytes()
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"600\" height=\"400\">\\n<path  d=\"M 0 0\nL 600 0\nL 600 35\nL 0 35\nL 0 0\" style=\"stroke-width:0;stroke:none;fill:rgba(240,240,240,1.0)\"/><path  d=\"M 0 0\nL 120 0\nL 120 35\nL 0 35\nL 0 0\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 120 0\nL 240 0\nL 240 35\nL 120 35\nL 120 0\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 240 0\nL 360 0\nL 360 35\nL 240 35\nL 240 0\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 360 0\nL 480 0\nL 480 35\nL 360 35\nL 360 0\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 480 0\nL 600 0\nL 600 35\nL 480 35\nL 480 0\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 0 35\nL 120 35\nL 120 90\nL 0 90\nL 0 35\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 120 35\nL 240 35\nL 240 90\nL 120 90\nL 120 35\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 240 35\nL 360 35\nL 360 90\nL 240 90\nL 240 35\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 360 35\nL 480 35\nL 480 90\nL 360 90\nL 360 35\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 0 90\nL 120 90\nL 120 145\nL 0 145\nL 0 90\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 120 90\nL 240 90\nL 240 145\nL 120 145\nL 120 90\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 240 90\nL 360 90\nL 360 145\nL 240 145\nL 240 90\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 360 90\nL 480 90\nL 480 145\nL 360 145\nL 360 90\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 0 145\nL 120 145\nL 120 200\nL 0 200\nL 0 145\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 120 145\nL 240 145\nL 240 200\nL 120 200\nL 120 145\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 240 145\nL 360 145\nL 360 200\nL 240 200\nL 240 145\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 360 145\nL 480 145\nL 480 200\nL 360 200\nL 360 145\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 480 145\nL 600 145\nL 600 200\nL 480 200\nL 480 145\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 480 35\nL 600 35\nL 600 145\nL 480 145\nL 480 35\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><text x=\"10\" y=\"22\" style=\"stroke-width:0;stroke:none;fill:rgba(98,105,118,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Name</text><text x=\"130\" y=\"22\" style=\"stroke-width:0;stroke:none;fill:rgba(98,105,118,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Age</text><text x=\"250\" y=\"22\" style=\"stroke-width:0;stroke:none;fill:rgba(98,105,118,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Address</text><text x=\"370\" y=\"22\" style=\"stroke-width:0;stroke:none;fill:rgba(98,105,118,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Tag</text><text x=\"490\" y=\"22\" style=\"stroke-width:0;stroke:none;fill:rgba(98,105,118,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Action</text><text x=\"10\" y=\"57\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">John Brown</text><text x=\"130\" y=\"57\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">32</text><text x=\"250\" y=\"57\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">New York No.</text><text x=\"250\" y=\"77\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">1 Lake Park</text><text x=\"370\" y=\"57\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">nice,</text><text x=\"370\" y=\"77\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">developer</text><text x=\"10\" y=\"112\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Jim Green</text><text x=\"130\" y=\"112\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">42</text><text x=\"250\" y=\"112\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">London No. 1</text><text x=\"250\" y=\"132\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Lake Park</text><text x=\"370\" y=\"112\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">wow</text><text x=\"10\" y=\"167\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Joe Black</text><text x=\"130\" y=\"167\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">32</text><text x=\"250\" y=\"167\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Sidney No. 1</text><text x=\"250\" y=\"187\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Lake Park</text><text x=\"370\" y=\"167\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">cool, teacher</text><text x=\"490\" y=\"167\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Send Mail</text><text x=\"490\" y=\"97\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Send Mail</text></svg>",
		},
		{
			render: func(p *Painter) ([]byte, error) {
				_, err := NewTableChart(p, TableChartOption{
					Header: []string{
						"Test",
					},
					Data: [][]string{
						{"0"},
						{"1"},
						{"1"},
						{"0"},
					},
					RowSpans: map[int][]CellSpan{
						0: {
							{
								RowFrom: 1,
								RowTo:   2,
							},
						},
					},
					StrokeColor:          drawing.ColorBlack,
					StrokeWidth:          1,
					DisableRowBackground: true,
				}).Render()
				if err != nil {
					return nil, err
				}
				return p.Bytes()
			},
			result: "",
		},
		{
			render: func(p *Painter) ([]byte, error) {
				_, err := NewTableChart(p, TableChartOption{
					Header: []string{
						"Test", "Test",
					},
					Data: [][]string{
						{"0", "42"},
						{"1", "42"},
						{"1", "42"},
						{"0", "42"},
					},
					RowSpans: map[int][]CellSpan{
						0: {
							{
								RowFrom: 1,
								RowTo:   2,
							},
						},
					},
					StrokeColor:          drawing.ColorBlack,
					StrokeWidth:          1,
					DisableRowBackground: true,
				}).Render()
				if err != nil {
					return nil, err
				}
				return p.Bytes()
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"600\" height=\"400\">\\n<path  d=\"M 0 0\nL 600 0\nL 600 35\nL 0 35\nL 0 0\" style=\"stroke-width:0;stroke:none;fill:rgba(240,240,240,1.0)\"/><path  d=\"M 0 0\nL 300 0\nL 300 35\nL 0 35\nL 0 0\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 300 0\nL 600 0\nL 600 35\nL 300 35\nL 300 0\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 0 35\nL 300 35\nL 300 70\nL 0 70\nL 0 35\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 300 35\nL 600 35\nL 600 70\nL 300 70\nL 300 35\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 300 70\nL 600 70\nL 600 105\nL 300 105\nL 300 70\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 300 105\nL 600 105\nL 600 140\nL 300 140\nL 300 105\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 0 140\nL 300 140\nL 300 175\nL 0 175\nL 0 140\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 300 140\nL 600 140\nL 600 175\nL 300 175\nL 300 140\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><path  d=\"M 0 70\nL 300 70\nL 300 140\nL 0 140\nL 0 70\" style=\"stroke-width:1;stroke:rgba(0,0,0,1.0);fill:none\"/><text x=\"10\" y=\"22\" style=\"stroke-width:0;stroke:none;fill:rgba(98,105,118,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Test</text><text x=\"310\" y=\"22\" style=\"stroke-width:0;stroke:none;fill:rgba(98,105,118,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Test</text><text x=\"10\" y=\"57\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">0</text><text x=\"310\" y=\"57\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">42</text><text x=\"310\" y=\"92\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">42</text><text x=\"310\" y=\"127\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">42</text><text x=\"10\" y=\"162\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">0</text><text x=\"310\" y=\"162\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">42</text><text x=\"10\" y=\"112\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">1</text></svg>",
		},
	}

	saveResult, _ := strconv.ParseBool(os.Getenv("SAVE_TEST_RESULT"))

	for i, tt := range tests {
		p, err := NewPainter(PainterOptions{
			Type:   ChartOutputSVG,
			Width:  600,
			Height: 400,
		}, PainterThemeOption(defaultTheme))
		assert.Nil(err)
		data, err := tt.render(p)
		assert.Nil(err)

		if saveResult {
			filename := fmt.Sprintf("table-%d.svg", i)
			err := os.WriteFile(filename, data, os.ModePerm)
			if err != nil {
				t.Logf("write %d result: %v", i, err)
			}
		}

		assert.Equal(tt.result, string(data))
	}
}
