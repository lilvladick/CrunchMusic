import SwiftUI

struct TabBarView: View {
    @State private var selectedPage: TabPage = .news
    @StateObject private var mainScreenViewModel = MainScreenViewModel()
    @StateObject private var breakingNewsViewModel = BreakingNewViewModel()
    @State private var isTabBarHidden: Bool = false
    
    init() {
        UITabBar.appearance().isHidden = true
    }
    
    var body: some View {
        ZStack(alignment: .bottom) {
            TabView(selection: $selectedPage) {
                MainScreenView(mainScreenViewModel: mainScreenViewModel, isTabBarHidden: $isTabBarHidden)
                    .tag(TabPage.news)
                BreakingNewsView(breakingNewsViewModel: breakingNewsViewModel, isTabBarHidden: $isTabBarHidden)
                    .tag(TabPage.breakingNews)
                SettingsView()
                    .tag(TabPage.settings)
            }
            
            CustomTabBar(selectedPage: $selectedPage)
                .padding(.horizontal)
                .opacity(isTabBarHidden ? 0 : 1)
        }
    }
}

#Preview {
    TabBarView()
}
