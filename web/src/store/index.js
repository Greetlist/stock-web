import Vuex from 'vuex'

const store = new Vuex.Store({
  state: {
    collapse: false
  },
  mutations: {
    changeCollapse (state, targetCollpaseState) {
      state.collapse = targetCollpaseState
    }
  }
})

export default store
