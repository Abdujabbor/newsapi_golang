package newsapi_golang

import (
	json "encoding/json"

	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson6601e8cdDecodeNewsapiGolang(in *jlexer.Lexer, out *Source) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "url":
			out.URL = string(in.String())
		case "category":
			out.Category = string(in.String())
		case "language":
			out.Language = string(in.String())
		case "country":
			out.Country = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeNewsapiGolang(out *jwriter.Writer, in Source) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"id\":")
	out.String(string(in.ID))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"name\":")
	out.String(string(in.Name))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"description\":")
	out.String(string(in.Description))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"url\":")
	out.String(string(in.URL))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"category\":")
	out.String(string(in.Category))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"language\":")
	out.String(string(in.Language))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"country\":")
	out.String(string(in.Country))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Source) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeNewsapiGolang(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Source) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeNewsapiGolang(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Source) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeNewsapiGolang(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Source) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeNewsapiGolang(l, v)
}
func easyjson6601e8cdDecodeNewsapiGolang1(in *jlexer.Lexer, out *Article) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "source":
			(out.Source).UnmarshalEasyJSON(in)
		case "author":
			out.Author = string(in.String())
		case "title":
			out.Title = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "url":
			out.URL = string(in.String())
		case "urlToImage":
			out.URLToImage = string(in.String())
		case "publishedAt":
			out.PublishedAt = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeNewsapiGolang1(out *jwriter.Writer, in Article) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"source\":")
	(in.Source).MarshalEasyJSON(out)
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"author\":")
	out.String(string(in.Author))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"title\":")
	out.String(string(in.Title))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"description\":")
	out.String(string(in.Description))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"url\":")
	out.String(string(in.URL))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"urlToImage\":")
	out.String(string(in.URLToImage))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"publishedAt\":")
	out.String(string(in.PublishedAt))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Article) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeNewsapiGolang1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Article) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeNewsapiGolang1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Article) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeNewsapiGolang1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Article) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeNewsapiGolang1(l, v)
}
func easyjson6601e8cdDecodeNewsapiGolang2(in *jlexer.Lexer, out *APIResponseSources) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "code":
			out.Code = string(in.String())
		case "message":
			out.Message = string(in.String())
		case "status":
			out.Status = string(in.String())
		case "sources":
			if in.IsNull() {
				in.Skip()
				out.Sources = nil
			} else {
				in.Delim('[')
				if out.Sources == nil {
					if !in.IsDelim(']') {
						out.Sources = make([]Source, 0, 1)
					} else {
						out.Sources = []Source{}
					}
				} else {
					out.Sources = (out.Sources)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Source
					(v1).UnmarshalEasyJSON(in)
					out.Sources = append(out.Sources, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeNewsapiGolang2(out *jwriter.Writer, in APIResponseSources) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"code\":")
	out.String(string(in.Code))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"message\":")
	out.String(string(in.Message))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"status\":")
	out.String(string(in.Status))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"sources\":")
	if in.Sources == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in.Sources {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v APIResponseSources) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeNewsapiGolang2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v APIResponseSources) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeNewsapiGolang2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *APIResponseSources) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeNewsapiGolang2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *APIResponseSources) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeNewsapiGolang2(l, v)
}
func easyjson6601e8cdDecodeNewsapiGolang3(in *jlexer.Lexer, out *APIResponseArticles) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "code":
			out.Code = string(in.String())
		case "message":
			out.Message = string(in.String())
		case "status":
			out.Status = string(in.String())
		case "totalResults":
			out.TotalResults = int(in.Int())
		case "articles":
			if in.IsNull() {
				in.Skip()
				out.Articles = nil
			} else {
				in.Delim('[')
				if out.Articles == nil {
					if !in.IsDelim(']') {
						out.Articles = make([]Article, 0, 1)
					} else {
						out.Articles = []Article{}
					}
				} else {
					out.Articles = (out.Articles)[:0]
				}
				for !in.IsDelim(']') {
					var v4 Article
					(v4).UnmarshalEasyJSON(in)
					out.Articles = append(out.Articles, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeNewsapiGolang3(out *jwriter.Writer, in APIResponseArticles) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"code\":")
	out.String(string(in.Code))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"message\":")
	out.String(string(in.Message))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"status\":")
	out.String(string(in.Status))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"totalResults\":")
	out.Int(int(in.TotalResults))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"articles\":")
	if in.Articles == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v5, v6 := range in.Articles {
			if v5 > 0 {
				out.RawByte(',')
			}
			(v6).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v APIResponseArticles) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeNewsapiGolang3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v APIResponseArticles) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeNewsapiGolang3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *APIResponseArticles) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeNewsapiGolang3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *APIResponseArticles) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeNewsapiGolang3(l, v)
}
