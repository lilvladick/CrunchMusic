import SwiftUI

struct TabBarView: View {
    @State private var selectedPage: TabPage = .music

    init() {
        UITabBar.appearance().isHidden = true
    }
    
    var body: some View {
        ZStack(alignment: .bottom) {
            TabView(selection: $selectedPage) {
                MainScreenView()
                    .tag(TabPage.music)
                BreakingNewsView()
                    .tag(TabPage.playlist)
                SettingsView()
                    .tag(TabPage.settings)
            }
            
            CustomTabBar(selectedPage: $selectedPage)
                .padding(.horizontal)
        }
    }
}

#Preview {
    TabBarView()
}
