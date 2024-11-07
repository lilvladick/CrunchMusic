import SwiftUI

struct MainScreenView: View {
    var body: some View {
        VStack {
            
        }
        .toolbar {
            ToolbarItem(placement: .topBarTrailing, content: {
                Button(action: {
                    //search action
                }, label: {
                    Image(systemName: "magnifyingglass")
                })
            })
        }
    }
}

#Preview {
    MainScreenView()
}
