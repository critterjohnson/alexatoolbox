package response

import (
	"testing"

	. "github.com/critterjohnson/go-ask/testgen"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewBuilder(t *testing.T) {
	Convey("When NewBuilder is called", t, func() {
		builder := NewBuilder()

		Convey("a new Builder is created", func() {
			So(builder, ShouldNotBeNil)
		})
	})
}

func TestWithAttributes(t *testing.T) {
	builder := NewBuilder()

	Convey("When WithAttributes is called", t, func() {
		attributes := make(map[string]interface{})

		Convey("with an empty map", func() {
			builder.WithAttributes(attributes)

			Convey("the builder should have no attributes", func() {
				So(builder.sessionAttributes, ShouldBeEmpty)
			})
		})

		Convey("with a populated map", func() {
			attributes[RandomString()] = RandomString()

			builder.WithAttributes(attributes)

			Convey("the builder receives the attributes", func() {
				So(builder.sessionAttributes, ShouldEqual, attributes)
			})
		})
	})
}

func TestAddAttribute(t *testing.T) {
	builder := NewBuilder()

	Convey("When AddAttribute is called", t, func() {
		attributes := make(map[string]interface{})

		key := RandomString()
		val := RandomString()
		attributes[key] = val

		builder.AddAttribute(key, val)
		Convey("the attribute should be added", func() {
			So(builder.sessionAttributes, ShouldResemble, attributes)
		})
	})
}

func TestAddAttributes(t *testing.T) {
	builder := NewBuilder()

	Convey("When AddAttributes is called", t, func() {
		attributes := make(map[string]interface{})

		Convey("with an empty map", func() {
			builder.AddAttributes(attributes)

			Convey("the builder shouldn't gain any attributes", func() {
				So(builder.sessionAttributes, ShouldBeEmpty)
			})
		})

		Convey("with a populated map", func() {
			attributes[RandomString()] = RandomString()
			builder.AddAttributes(attributes)

			Convey("the builder should take on the new attributes", func() {
				So(builder.sessionAttributes, ShouldResemble, attributes)
			})
		})
	})
}

func TestWithOutputSpeech(t *testing.T) {
	builder := NewBuilder()

	Convey("When WithOutputSpeech is called", t, func() {
		Convey("with an empty object", func() {
			outputSpeech := OutputSpeech{}
			builder.WithOutputSpeech(outputSpeech)

			Convey("the builder's outputSpeech is empty", func() {
				So(*builder.outputSpeech, ShouldResemble, outputSpeech)
			})
		})

		Convey("with a populated object", func() {
			outputSpeech := OutputSpeech{
				Type: RandomString(),
				Text: RandomString(),
				Ssml: RandomString(),
			}
			builder.WithOutputSpeech(outputSpeech)

			Convey("the builder should have the OutputSpeech object", func() {
				So(*builder.outputSpeech, ShouldResemble, outputSpeech)
			})
		})
	})
}

func TestWithTextOutputSpeech(t *testing.T) {
	builder := NewBuilder()

	Convey("When WithTextOutputSpeech is called", t, func() {
		text := RandomString()
		builder.WithTextOutputSpeech(text)

		Convey("and the OutputSpeech object already exists", func() {
			playBehavior := RandomString()
			builder.OutputSpeechPlayBehavior(playBehavior)

			Convey("it should not overrite the existing OutputSpeech object", func() {
				So(builder.outputSpeech.PlayBehavior, ShouldEqual, playBehavior)
			})
		})

		Convey("it should create the OutputSpeech object", func() {
			So(builder.outputSpeech, ShouldNotBeNil)
		})

		Convey("the OutputSpeech type should be PlainText", func() {
			So(builder.outputSpeech.Type, ShouldEqual, "PlainText")
		})

		Convey("it should have the right text", func() {
			So(builder.outputSpeech.Text, ShouldEqual, text)
		})
	})
}

func TestWithSsmlOutputSpeech(t *testing.T) {
	builder := NewBuilder()

	Convey("When WithSsmlOutputSpeech is called", t, func() {
		text := RandomString()
		builder.WithSsmlOutputSpeech(text)

		Convey("and the OutputSpeech object already exists", func() {
			playBehavior := RandomString()
			builder.OutputSpeechPlayBehavior(playBehavior)

			Convey("it should not overrite the existing OutputSpeech object", func() {
				So(builder.outputSpeech.PlayBehavior, ShouldEqual, playBehavior)
			})
		})

		Convey("it should create the OutputSpeech object", func() {
			So(builder.outputSpeech, ShouldNotBeNil)
		})

		Convey("the OutputSpeech type should be ssml", func() {
			So(builder.outputSpeech.Type, ShouldEqual, "ssml")
		})

		Convey("it should have the right text", func() {
			So(builder.outputSpeech.Ssml, ShouldEqual, text)
		})
	})
}

func TestOutputSpeechPlayBehavior(t *testing.T) {
	builder := NewBuilder()

	Convey("When OutputSpeechPlayBehavior is called", t, func() {
		behavior := RandomString()

		Convey("and the OutputSpeech object already exists", func() {
			outputSpeech := OutputSpeech{
				Type: RandomString(),
				Text: RandomString(),
				Ssml: RandomString(),
			}
			builder.WithOutputSpeech(outputSpeech)
			builder.OutputSpeechPlayBehavior(behavior)

			Convey("it should set the PlayBehavior", func() {
				So(builder.outputSpeech.PlayBehavior, ShouldEqual, behavior)
			})

			Convey("without disrupting the original object", func() {
				original := [3]string{outputSpeech.Type, outputSpeech.Text, outputSpeech.Ssml}
				new := [3]string{builder.outputSpeech.Type, builder.outputSpeech.Text, builder.outputSpeech.Ssml}
				So(original, ShouldEqual, new)
			})
		})
		Convey("and the OutputSpeech object doesn't exist", func() {
			builder.OutputSpeechPlayBehavior(behavior)

			Convey("it should create a new OutputSpeech object", func() {
				So(builder.outputSpeech, ShouldNotBeNil)
			})

			Convey("it should set the PlayBehavior", func() {
				So(builder.outputSpeech.PlayBehavior, ShouldEqual, behavior)
			})
		})
	})
}

func TestWithCard(t *testing.T) {
	builder := NewBuilder()

	Convey("When WithCard is called", t, func() {
		Convey("with an empty object", func() {
			card := Card{}
			builder.WithCard(card)

			Convey("the builder's card is empty", func() {
				So(*builder.card, ShouldResemble, card)
			})
		})

		Convey("with a populated object", func() {
			card := Card{
				Type: RandomString(),
				Text: RandomString(),
			}
			builder.WithCard(card)

			Convey("the builder should have the Card object", func() {
				So(*builder.card, ShouldResemble, card)
			})
		})
	})
}

func TestWithReprompt(t *testing.T) {
	builder := NewBuilder()

	Convey("When WithReprompt is called", t, func() {
		Convey("with an empty object", func() {
			outputSpeech := OutputSpeech{}
			reprompt := Reprompt{
				OutputSpeech: &outputSpeech,
			}
			builder.WithReprompt(reprompt)

			Convey("the Reprompt's outputSpeech is empty", func() {
				So(*builder.reprompt, ShouldResemble, reprompt)
			})
		})

		Convey("with a populated object", func() {
			outputSpeech := OutputSpeech{
				Type: RandomString(),
				Text: RandomString(),
				Ssml: RandomString(),
			}
			reprompt := Reprompt{
				OutputSpeech: &outputSpeech,
			}
			builder.WithReprompt(reprompt)

			Convey("the Reprompt should have the OutputSpeech object", func() {
				So(*builder.reprompt, ShouldResemble, reprompt)
			})
		})
	})
}

func TestWithTextReprompt(t *testing.T) {
	builder := NewBuilder()

	Convey("When WithTextReprompt is called", t, func() {
		text := RandomString()
		builder.WithTextReprompt(text)

		Convey("and the reprompt.OutputSpeech object already exists", func() {
			playBehavior := RandomString()
			builder.RepromptPlayBehavior(playBehavior)

			Convey("it should not overrite the existing OutputSpeech object", func() {
				So(builder.reprompt.OutputSpeech.PlayBehavior, ShouldEqual, playBehavior)
			})
		})

		Convey("it should create the Reprompt object", func() {
			So(builder.reprompt, ShouldNotBeNil)
		})

		Convey("the OutputSpeech type should be PlainText", func() {
			So(builder.reprompt.OutputSpeech.Type, ShouldEqual, "PlainText")
		})

		Convey("it should have the right text", func() {
			So(builder.reprompt.OutputSpeech.Text, ShouldEqual, text)
		})
	})
}

func TestWithSsmlReprompt(t *testing.T) {
	builder := NewBuilder()

	Convey("When WithSsmlReprompt is called", t, func() {
		ssml := RandomString()
		builder.WithSsmlReprompt(ssml)

		Convey("and the reprompt.OutputSpeech object already exists", func() {
			playBehavior := RandomString()
			builder.RepromptPlayBehavior(playBehavior)

			Convey("it should not overrite the existing OutputSpeech object", func() {
				So(builder.reprompt.OutputSpeech.PlayBehavior, ShouldEqual, playBehavior)
			})
		})

		Convey("it should create the Reprompt object", func() {
			So(builder.reprompt, ShouldNotBeNil)
		})

		Convey("the OutputSpeech type should be ssml", func() {
			So(builder.reprompt.OutputSpeech.Type, ShouldEqual, "ssml")
		})

		Convey("it should have the right ssml", func() {
			So(builder.reprompt.OutputSpeech.Ssml, ShouldEqual, ssml)
		})
	})
}

func TestRepromptPlayBehavior(t *testing.T) {
	builder := NewBuilder()

	Convey("When RepromptPlayBehavior is called", t, func() {
		behavior := RandomString()

		Convey("and the Reprompt object already exists", func() {
			outputSpeech := OutputSpeech{
				Type: RandomString(),
				Text: RandomString(),
				Ssml: RandomString(),
			}
			reprompt := Reprompt{
				OutputSpeech: &outputSpeech,
			}
			builder.WithReprompt(reprompt)
			builder.RepromptPlayBehavior(behavior)

			Convey("it should set the PlayBehavior", func() {
				So(builder.reprompt.OutputSpeech.PlayBehavior, ShouldEqual, behavior)
			})

			Convey("without disrupting the original object", func() {
				original := [3]string{outputSpeech.Type, outputSpeech.Text, outputSpeech.Ssml}
				new := [3]string{builder.reprompt.OutputSpeech.Type,
					builder.reprompt.OutputSpeech.Text,
					builder.reprompt.OutputSpeech.Ssml}
				So(original, ShouldEqual, new)
			})
		})
		Convey("and the OutputSpeech object doesn't exist", func() {
			builder.RepromptPlayBehavior(behavior)

			Convey("it should create a new Reprompt object", func() {
				So(builder.reprompt, ShouldNotBeNil)
			})

			Convey("it should set the PlayBehavior", func() {
				So(builder.reprompt.OutputSpeech.PlayBehavior, ShouldEqual, behavior)
			})
		})
	})
}

func TestWithDirectives(t *testing.T) {
	Convey("When WithDirectives is called", t, func() {
		builder := NewBuilder()
		directives := []interface{}{RandomString(), RandomString(), RandomString()}

		builder.WithDirectives(directives)

		Convey("it should set the directives", func() {
			So(builder.directives, ShouldResemble, directives)
		})
	})
}

func TestShouldEndSession(t *testing.T) {
	Convey("When ShouldEndSession is called", t, func() {
		builder := NewBuilder()

		builder.ShouldEndSession(false)

		Convey("it sets the ShouldEndSession boolean", func() {
			So(builder.shouldEndSession, ShouldBeFalse)
		})
	})
}

func TestBuild(t *testing.T) {
	Convey("When Build is called", t, func() {
		builder := NewBuilder()

		response := builder.Build()

		Convey("it should return the Response object", func() {
			So(response, ShouldHaveSameTypeAs, Response{})
		})
	})
}
