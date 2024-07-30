package components

import carbon "github.com/TheWozard/go-carbonated"

type Watcher struct {
	Watch   func(msg *carbon.Msg)
	Content carbon.Component
}

func (w Watcher) ComponentUpdate(msg *carbon.Msg, cmd *carbon.Cmd) carbon.Component {
	w.Watch(msg)
	w.Content = w.Content.ComponentUpdate(msg, cmd)
	return w
}

func (w Watcher) View() string {
	return w.Content.View()
}
