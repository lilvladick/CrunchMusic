import Foundation

class Formatter {
    private let timestamp: String
    private var dateComponents: DateComponents?
    
    init(timestamp: String) {
        self.timestamp = timestamp
        self.dateComponents = Formatter.parseDateComponents(timestamp)
    }
    
    private static func parseDateComponents(_ timestamp: String) -> DateComponents? {
        let isoFormatter = ISO8601DateFormatter()
        isoFormatter.formatOptions = [.withInternetDateTime, .withFractionalSeconds]
        
        guard let date = isoFormatter.date(from: timestamp) else { return nil }
        
        let calendar = Calendar.current
        return calendar.dateComponents([.year, .month, .day, .hour, .minute, .second], from: date)
    }
    
    func getFormattedDate() -> String {
        guard let components = dateComponents else { return "N/A" }
        
        guard let year = components.year, let month = components.month, let day = components.day else { return "N/A" }
        
        return String(format: "%02d.%02d.%04d", day, month, year)
    }
    
    func getFormattedTime() -> String {
        guard let components = dateComponents else { return "N/A" }
        
        guard let hour = components.hour, let minute = components.minute else { return "N/A" }
        
        return String(format: "%02d:%02d", hour, minute)
    }
}
