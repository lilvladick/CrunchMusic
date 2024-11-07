import SwiftUI

struct CustomTabBar: View {
    @Binding var selectedPage: TabPage
    
    var body: some View {
        HStack {
            ForEach(TabPage.allCases, id: \.self) { page in
                TabBarItem(page: page, selectedPage: $selectedPage)
            }
        }
        .frame(maxWidth: .infinity)
    }
}
