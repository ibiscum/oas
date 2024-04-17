// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/ogen-go/ogen/uri"
)

func (s *Server) cutPrefix(path string) (string, bool) {
	prefix := s.cfg.Prefix
	if prefix == "" {
		return path, true
	}
	if !strings.HasPrefix(path, prefix) {
		// Prefix doesn't match.
		return "", false
	}
	// Cut prefix from the path.
	return strings.TrimPrefix(path, prefix), true
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	elemIsEscaped := false
	if rawPath := r.URL.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
			elemIsEscaped = strings.ContainsRune(elem, '%')
		}
	}

	elem, ok := s.cutPrefix(elem)
	if !ok || len(elem) == 0 {
		s.notFound(w, r)
		return
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			origElem := elem
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'a': // Prefix: "anything/"
				origElem := elem
				if l := len("anything/"); len(elem) >= l && elem[0:l] == "anything/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'a': // Prefix: "a"
					origElem := elem
					if l := len("a"); len(elem) >= l && elem[0:l] == "a" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'l': // Prefix: "ll-of-object"
						origElem := elem
						if l := len("ll-of-object"); len(elem) >= l && elem[0:l] == "ll-of-object" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleAnythingAllOfObjectRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}

						elem = origElem
					case 'n': // Prefix: "ny-of-primitive"
						origElem := elem
						if l := len("ny-of-primitive"); len(elem) >= l && elem[0:l] == "ny-of-primitive" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleAnythingAnyOfPrimitiveRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}

						elem = origElem
					}

					elem = origElem
				case 'n': // Prefix: "nested-one-of-object-ref"
					origElem := elem
					if l := len("nested-one-of-object-ref"); len(elem) >= l && elem[0:l] == "nested-one-of-object-ref" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleAnythingNestedOneOfObjectRefRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}

					elem = origElem
				case 'o': // Prefix: "one-of-"
					origElem := elem
					if l := len("one-of-"); len(elem) >= l && elem[0:l] == "one-of-" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'o': // Prefix: "object"
						origElem := elem
						if l := len("object"); len(elem) >= l && elem[0:l] == "object" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch r.Method {
							case "POST":
								s.handleAnythingOneOfObjectRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}
						switch elem[0] {
						case '-': // Prefix: "-ref"
							origElem := elem
							if l := len("-ref"); len(elem) >= l && elem[0:l] == "-ref" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "POST":
									s.handleAnythingOneOfObjectRefRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "POST")
								}

								return
							}

							elem = origElem
						}

						elem = origElem
					case 'p': // Prefix: "primitive"
						origElem := elem
						if l := len("primitive"); len(elem) >= l && elem[0:l] == "primitive" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleAnythingOneOfPrimitiveRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}

						elem = origElem
					}

					elem = origElem
				}

				elem = origElem
			case 'p': // Prefix: "pets"
				origElem := elem
				if l := len("pets"); len(elem) >= l && elem[0:l] == "pets" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "PATCH":
						s.handleGetPetsRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "PATCH")
					}

					return
				}

				elem = origElem
			}

			elem = origElem
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name        string
	summary     string
	operationID string
	pathPattern string
	count       int
	args        [0]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// Summary returns OpenAPI summary.
func (r Route) Summary() string {
	return r.summary
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
}

// PathPattern returns OpenAPI path.
func (r Route) PathPattern() string {
	return r.pathPattern
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
//
// Note: this method does not unescape path or handle reserved characters in path properly. Use FindPath instead.
func (s *Server) FindRoute(method, path string) (Route, bool) {
	return s.FindPath(method, &url.URL{Path: path})
}

// FindPath finds Route for given method and URL.
func (s *Server) FindPath(method string, u *url.URL) (r Route, _ bool) {
	var (
		elem = u.Path
		args = r.args
	)
	if rawPath := u.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
		defer func() {
			for i, arg := range r.args[:r.count] {
				if unescaped, err := url.PathUnescape(arg); err == nil {
					r.args[i] = unescaped
				}
			}
		}()
	}

	elem, ok := s.cutPrefix(elem)
	if !ok {
		return r, false
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			origElem := elem
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'a': // Prefix: "anything/"
				origElem := elem
				if l := len("anything/"); len(elem) >= l && elem[0:l] == "anything/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'a': // Prefix: "a"
					origElem := elem
					if l := len("a"); len(elem) >= l && elem[0:l] == "a" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'l': // Prefix: "ll-of-object"
						origElem := elem
						if l := len("ll-of-object"); len(elem) >= l && elem[0:l] == "ll-of-object" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "POST":
								// Leaf: AnythingAllOfObject
								r.name = "AnythingAllOfObject"
								r.summary = "allOf with listed objects"
								r.operationID = "anything-all-of-object"
								r.pathPattern = "/anything/all-of-object"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

						elem = origElem
					case 'n': // Prefix: "ny-of-primitive"
						origElem := elem
						if l := len("ny-of-primitive"); len(elem) >= l && elem[0:l] == "ny-of-primitive" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "POST":
								// Leaf: AnythingAnyOfPrimitive
								r.name = "AnythingAnyOfPrimitive"
								r.summary = "anyOf primitive"
								r.operationID = "anything-any-of-primitive"
								r.pathPattern = "/anything/any-of-primitive"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

						elem = origElem
					}

					elem = origElem
				case 'n': // Prefix: "nested-one-of-object-ref"
					origElem := elem
					if l := len("nested-one-of-object-ref"); len(elem) >= l && elem[0:l] == "nested-one-of-object-ref" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "POST":
							// Leaf: AnythingNestedOneOfObjectRef
							r.name = "AnythingNestedOneOfObjectRef"
							r.summary = "nested oneOf object with $ref pointers"
							r.operationID = "anything-nested-one-of-object-ref"
							r.pathPattern = "/anything/nested-one-of-object-ref"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

					elem = origElem
				case 'o': // Prefix: "one-of-"
					origElem := elem
					if l := len("one-of-"); len(elem) >= l && elem[0:l] == "one-of-" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'o': // Prefix: "object"
						origElem := elem
						if l := len("object"); len(elem) >= l && elem[0:l] == "object" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "POST":
								r.name = "AnythingOneOfObject"
								r.summary = "oneOf object"
								r.operationID = "anything-one-of-object"
								r.pathPattern = "/anything/one-of-object"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
						switch elem[0] {
						case '-': // Prefix: "-ref"
							origElem := elem
							if l := len("-ref"); len(elem) >= l && elem[0:l] == "-ref" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "POST":
									// Leaf: AnythingOneOfObjectRef
									r.name = "AnythingOneOfObjectRef"
									r.summary = "oneOf object with $ref pointers"
									r.operationID = "anything-one-of-object-ref"
									r.pathPattern = "/anything/one-of-object-ref"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

							elem = origElem
						}

						elem = origElem
					case 'p': // Prefix: "primitive"
						origElem := elem
						if l := len("primitive"); len(elem) >= l && elem[0:l] == "primitive" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "POST":
								// Leaf: AnythingOneOfPrimitive
								r.name = "AnythingOneOfPrimitive"
								r.summary = "oneOf primitive"
								r.operationID = "anything-one-of-primitive"
								r.pathPattern = "/anything/one-of-primitive"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

						elem = origElem
					}

					elem = origElem
				}

				elem = origElem
			case 'p': // Prefix: "pets"
				origElem := elem
				if l := len("pets"); len(elem) >= l && elem[0:l] == "pets" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "PATCH":
						// Leaf: GetPets
						r.name = "GetPets"
						r.summary = "oneOf request with a nested allOf"
						r.operationID = "get_pets"
						r.pathPattern = "/pets"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}

				elem = origElem
			}

			elem = origElem
		}
	}
	return r, false
}