package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected *node
	}{
		{
			"simple array string parsing",
			"[a,b,c]",
			newNode("", []*node{
				newNode("a", []*node{}),
				newNode("b", []*node{}),
				newNode("c", []*node{}),
			}),
		},
		{
			"nested array string parsing 3 level",
			"[a[aa[aaa],ab,ac],b,c[ca,cb,cc[cca]]]",
			newNode("", []*node{
				newNode("a", []*node{
					newNode("aa", []*node{
						newNode("aaa", []*node{}),
					}),
					newNode("ab", []*node{}),
					newNode("ac", []*node{}),
				}),
				newNode("b", []*node{}),
				newNode("c", []*node{
					newNode("ca", []*node{}),
					newNode("cb", []*node{}),
					newNode("cc", []*node{
						newNode("cca", []*node{}),
					}),
				}),
			}),
		},
		{
			"nested array string parsing 6 level",
			"[a[ab[abb[aaaa[aaaaa[aaaaaa]]],abc],ac]]",
			newNode("", []*node{
				newNode("a", []*node{
					newNode("ab", []*node{
						newNode("abb", []*node{
							newNode("aaaa", []*node{
								newNode("aaaaa", []*node{
									newNode("aaaaaa", []*node{}),
								}),
							}),
						}),
						newNode("abc", []*node{}),
					}),
					newNode("ac", []*node{}),
				}),
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n, err := parse(tt.input)
			require.Nil(t, err)
			require.Equal(t, tt.expected, n)
		})
	}
}
