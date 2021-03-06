/*
 * This file is part of com.pharbers.DL.
 *
 * com.pharbers.DL is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * com.pharbers.DL is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with Foobar.  If not, see <https://www.gnu.org/licenses/>.
 */
package PhFormat

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAddColFormat_Exec(t *testing.T) {
	Convey("增加常数列", t, func() {
		data := []map[string]interface{}{
			{
				"firstname": "A",
				"lastname": "a",
				"age": 11,
				"factor": 3,
			},
			{
				"firstname": "B",
				"lastname": "b",
				"age": 22,
				"factor": 5,
			},
			{
				"firstname": "C",
				"lastname": "c",
				"age": 33,
				"factor": 7,
			},
		}

		addCols := []interface{}{
			map[string]interface{}{
				"name": "newCol1",
				"value": "value1",
			},
			map[string]interface{}{
				"name": "newCol2",
				"value": "value2",
			},
		}
		result, err := AddColFormat{}.Exec(addCols)(data)

		So(err, ShouldBeNil)
		So(result, ShouldNotBeNil)

		row := result.([]map[string]interface{})[0]
		var keys = make([]string, 0)
		for k, _ := range row {
			keys = append(keys, k)
		}
		So(len(keys), ShouldEqual, 5)
	})

	Convey("增加引用列", t, func() {
		data := []map[string]interface{}{
			{
				"firstname": "A",
				"lastname": "a",
				"age": 11,
				"factor": 3,
			},
			{
				"firstname": "B",
				"lastname": "b",
				"age": 22,
				"factor": 5,
			},
			{
				"firstname": "C",
				"lastname": "c",
				"age": 33,
				"factor": 7,
			},
		}

		addCols := []interface{}{
			map[string]interface{}{
				"name": "newCol1",
				"value": []interface{}{"$", "age"},
			},
			map[string]interface{}{
				"name": "newCol2",
				"value": []interface{}{"$", "factor"},
			},
		}
		result, err := AddColFormat{}.Exec(addCols)(data)

		So(err, ShouldBeNil)
		So(result, ShouldNotBeNil)

		row := result.([]map[string]interface{})[0]
		var keys = make([]string, 0)
		for k, _ := range row {
			keys = append(keys, k)
		}
		So(len(keys), ShouldEqual, 6)
	})

	Convey("增加计算列", t, func() {
		data := []map[string]interface{}{
			{
				"firstname": "A",
				"lastname": "a",
				"age": 11,
				"factor": 3,
			},
			{
				"firstname": "B",
				"lastname": "b",
				"age": 22,
				"factor": 5,
			},
			{
				"firstname": "C",
				"lastname": "c",
				"age": 33,
				"factor": 7,
			},
		}

		addCols := []interface{}{
			map[string]interface{}{
				"name": "newCol1",
				"value": []interface{}{"*", "age", "factor"},
			},
			map[string]interface{}{
				"name": "newCol2",
				"value": []interface{}{"/", "age", "factor"},
			},
		}
		result, err := AddColFormat{}.Exec(addCols)(data)

		So(err, ShouldBeNil)
		So(result, ShouldNotBeNil)

		row := result.([]map[string]interface{})[0]
		var keys = make([]string, 0)
		for k, _ := range row {
			keys = append(keys, k)
		}
		So(len(keys), ShouldEqual, 6)
	})

	Convey("增加嵌套计算列", t, func() {
		data := []map[string]interface{}{
			{
				"firstname": "A",
				"lastname": "a",
				"age": 11,
				"factor": 3,
			},
			{
				"firstname": "B",
				"lastname": "b",
				"age": 22,
				"factor": 5,
			},
			{
				"firstname": "C",
				"lastname": "c",
				"age": 33,
				"factor": 7,
			},
		}

		addCols := []interface{}{
			map[string]interface{}{
				"name": "newCol1",
				"value": []interface{}{"*", []interface{}{"+", "age", "age"}, "factor"},
			},
			map[string]interface{}{
				"name": "newCol2",
				"value": []interface{}{"/", []interface{}{"+", "age", []interface{}{"=", 1}}, "factor"},
			},
		}
		result, err := AddColFormat{}.Exec(addCols)(data)

		So(err, ShouldBeNil)
		So(result, ShouldNotBeNil)

		row := result.([]map[string]interface{})[0]
		var keys = make([]string, 0)
		for k, _ := range row {
			keys = append(keys, k)
		}
		So(len(keys), ShouldEqual, 6)
	})
}
