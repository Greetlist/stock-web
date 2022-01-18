import Vuex from 'vuex'

const store = new Vuex.Store({
  state: {
    collapse: false,
    selectAlgo: ''
  },
  mutations: {
    changeCollapse (state, targetCollpaseState) {
      state.collapse = targetCollpaseState
    },
    changeSelectAlgo (state, algo) {
      state.selectAlgo = algo
    }
  }
})

export default store
