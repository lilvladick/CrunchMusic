import SwiftUI

struct TabBarItem: View {
    var page: TabPage
    @Binding var selectedPage: TabPage
    
    private var isSelected: Bool {
        selectedPage == page
    }
    
    var body: some View {
        Button(action: {
            withAnimation(.snappy) {
                selectedPage = page
            }
        }) {
            HStack {
                Image(systemName: page.icon)
                    .foregroundColor(isSelected ? .blue : .gray)
                if isSelected {
                    Text(page.label)
                        .foregroundColor(.blue)
                        .fontWeight(.bold)
                }
            }
            .padding(.vertical, 8)
            .frame(maxWidth: .infinity)
            .background(isSelected ? Color.white : .clear)
            .clipShape(Capsule())
            .font(isSelected ? .callout : .title3)
        }
    }
}
