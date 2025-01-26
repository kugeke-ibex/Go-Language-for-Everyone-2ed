const app = new Vue({
  el: '#app',
  data: {
		comments: [],
		name: "",
		text: "",
	},

  created: function() {
    this.updateComments();
  },

  methods: {
    addComment: () => {
      const payload = {
        'name': app.name,
        'text': app.text,
      }

      axios.post("/api/comments", payload).then(() => {
      }).then(() => {
        app.name = ""
        app.text = ""
        app.updateComments()
      }).catch((err) => {
        alert(err.response.data)
      })
    },

    updateComments: () => {
      axios.get("/api/comments")
      .then((response) => app.comments = response.data || [])
      .catch((error) => console.log(error));
    },
  }
});
